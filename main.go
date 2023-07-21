package main

import (
	"fmt"
	"os"
	"time"

	"github.com/New-Earth-Lab/bgapi-go/bgapi2"
)

func BufferHandler(dataStream *bgapi2.DataStream, buffer *bgapi2.Buffer, userObject any) {
	if val, _ := buffer.GetIsIncomplete(); val {
		return
	} else {
		frameId, _ := buffer.GetFrameId()
		buf, _ := buffer.GetMemImageBuffer()
		fmt.Printf("Image %d received in buffer, length: %d\n", frameId, buf.Capacity())

	}
	dataStream.QueueBuffer(buffer)
}

func PrintDeviceInfo(device *bgapi2.Device) error {
	model, err := device.GetModel()
	if err != nil {
		return err
	}

	serialNumber, err := device.GetSerialNumber()
	if err != nil {
		return err
	}

	fmt.Printf("Model: %s(%s)\n", model, serialNumber)

	return nil
}

type Config struct {
	System       string
	Camera       string
	SerialNumber string
}

func ReleaseResources(device *bgapi2.Device) error {
	err := device.Close()
	if err != nil {
		return nil
	}

	iface, err := device.GetParent()
	if err != nil {
		return nil
	}

	err = iface.Close()
	if err != nil {
		return nil
	}

	system, err := iface.GetParent()
	if err != nil {
		return nil
	}

	err = system.Close()
	if err != nil {
		return nil
	}

	err = system.ReleaseSystem()
	if err != nil {
		return nil
	}

	return nil
}

// Get Device based on system:model:serial
func GetDevice(config Config) (*bgapi2.Device, error) {
	systems, err := bgapi2.GetSystems()
	if err != nil {
		return nil, err
	}

	var system *bgapi2.System
	for _, sys := range systems {
		fileName, err := sys.GetFileName()
		if err != nil {
			return nil, err
		}

		if fileName == config.System {
			system = sys
			break
		}
	}

	if system == nil {
		return nil, fmt.Errorf("could not find system")
	}

	err = system.Open()
	if err != nil {
		return nil, err
	}
	defer func(e *error) {
		if *e != nil {
			system.Close()
			system.ReleaseSystem()
		}
	}(&err)

	interfaces, err := system.GetInterfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		err = i.Open()
		if err != nil {
			return nil, err
		}
		defer func(e *error) {
			if *e != nil {
				i.Close()
			}
		}(&err)

		devices, err := i.GetDevices()
		if err != nil {
			return nil, err
		}

		for _, device := range devices {
			err = device.Open()
			if err != nil {
				return nil, err
			}
			defer func(e *error) {
				if *e != nil {
					device.Close()
				}
			}(&err)

			model, err := device.GetModel()
			if err != nil {
				return nil, err
			}

			serialNumber, err := device.GetSerialNumber()
			if err != nil {
				return nil, err
			}

			if model == config.Camera && serialNumber == config.SerialNumber {
				// Found it
				return device, nil
			}

			err = device.Close()
			if err != nil {
				return nil, err
			}
		}

		err = i.Close()
		if err != nil {
			return nil, err
		}
	}

	// Have to set err so cleanup code is called properly
	err = fmt.Errorf("could not find device")

	return nil, err
}

func AddBuffers(dataStream *bgapi2.DataStream, bufferCount int) error {

	for indexBuffers := 0; indexBuffers < bufferCount; indexBuffers++ {
		buffer, err := bgapi2.CreateBuffer()
		if err != nil {
			return err
		}

		err = dataStream.AnnounceBuffer(buffer)
		if err != nil {
			return err
		}

		dataStream.QueueBuffer(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func StartStreaming(device *bgapi2.Device) (*bgapi2.DataStream, error) {
	// Just assuming stream 0
	stream, err := device.GetDataStream(0)
	if err != nil {
		return nil, err
	}

	err = stream.Open()
	if err != nil {
		return nil, err
	}

	const bufferCount = 4

	err = AddBuffers(stream, bufferCount)
	if err != nil {
		return nil, err
	}

	err = stream.SetNewBufferEventMode(bgapi2.EventMode_EventHandler)
	if err != nil {
		return nil, err
	}

	err = stream.RegisterNewBufferEventHandler(BufferHandler, nil)
	if err != nil {
		return nil, err
	}

	err = stream.StartAcquisitionContinuous()
	if err != nil {
		return nil, err
	}

	node, err := device.GetRemoteNode("AcquisitionStart")
	if err != nil {
		return nil, err
	}

	err = node.Execute()
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func CaptureFrames(stream *bgapi2.DataStream) error {
	numDelivered := uint64(0)
	const numToCapture = 10

	var err error
	for numDelivered < numToCapture {
		numDelivered, err = stream.GetNumDelivered()
		if err != nil {
			return err
		}

		time.Sleep(10 * time.Millisecond)
	}

	return nil
}

func StopStreaming(device *bgapi2.Device, stream *bgapi2.DataStream) error {
	node, err := device.GetRemoteNode("AcquisitionAbort")
	if err != nil {
		return err
	}

	err = node.Execute()
	if err != nil {
		return err
	}

	node, err = device.GetRemoteNode("AcquisitionStop")
	if err != nil {
		return err
	}

	err = node.Execute()
	if err != nil {
		return err
	}

	err = stream.StopAcquisition()
	if err != nil {
		return err
	}

	isGrabbing := true
	for isGrabbing {
		isGrabbing, err = stream.IsGrabbing()
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	err = stream.SetNewBufferEventMode(bgapi2.EventMode_Polling)
	if err != nil {
		return err
	}

	// Remove buffers
	err = stream.Close()
	if err != nil {
		return err
	}

	return nil
}

func setDeviceParameter(device *bgapi2.Device) error {
	node, err := device.GetRemoteNode("TriggerMode")
	if err != nil {
		return err
	}
	err = node.SetString("Off")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// config := Config{
	// 	System:       "libbgapi2_usb.cti",
	// 	Camera:       "acA640-90um",
	// 	SerialNumber: "22367637",
	// }

	config := Config{
		System:       "libbgapi2_gige.cti",
		Camera:       "Goldeye G-008 Cool (4068580)",
		SerialNumber: "08-406858000765",
	}

	device, err := GetDevice(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer ReleaseResources(device)

	err = PrintDeviceInfo(device)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = setDeviceParameter(device)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	stream, err := StartStreaming(device)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = CaptureFrames(stream)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = StopStreaming(device, stream)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
