# QingCloud Service Usage Guide

Import and initialize QingCloud service with a config, and you are ready to use the initialized service.

Each API function take a Input struct and return an Output struct. The Input struct consists of request params, request headers and request elements, and the Output holds the HTTP status code, response headers, response elements and error (if error occurred).

``` go
import (
	"github.com/yunify/qingcloud-sdk-go/service"
)
```

### Code Snippet

Initialize the QingCloud service with a configuration

``` go
qcService, _ := qingcloud.Init(configuration)
```

Initialize the instance service in a zone

``` go
pek3aInstance, _ := qcService.Instance("pek3a")
```

DescribeInstances

``` go
iOutput, _ := pek3aInstance.DescribeInstances(
	&qingcloud.DescribeInstancesInput{
		Instances: []string{"i-xxxxxxxx"},
	},
)

// Print the return code.
fmt.Println(iOutput.RetCode)

// Print the first instance ID.
fmt.Println(iOutput.InstanceSet[0].InstanceID)
```

RunInstances

``` go
iOutput, _ := pek3aInstance.RunInstances(
	&qingcloud.RunInstancesInput{
		ImageID:      "centos7x64d",
		CPU:          1,
		Memory:       1024,
		LoginMode:    "keypair",
		LoginKeyPair: "kp-xxxxxxxx",
	},
)

// Print the return code.
fmt.Println(iOutput.RetCode)

// Print the job ID.
fmt.Println(iOutput.JobID)
```

Initialize the volume service in a zone

``` go
pek3aVolume, _ := qcService.Volume("pek3a")
```

DescribeVolumes

``` go
volOutput, _ := pek3aVolume.DescribeVolumes(&qingcloud.DescribeVolumesInput{
	Volumes: []string{"vol-xxxxxxxx"},
})

// Print the return code.
fmt.Println(volOutput.RetCode)

// Print the first volume name.
fmt.Println(volOutput.VolumeSet[0].VolumeName)
```

CreateVolumes

``` go
volOutput, _ := pek3aVolume.CreateVolumes(
	&qingcloud.CreateVolumesInput{
		Size:  10,
		Count: 2,
	},
)

// Print the return code.
fmt.Println(volOutput.RetCode)

// Print the job ID.
fmt.Println(volOutput.JobID)
```

