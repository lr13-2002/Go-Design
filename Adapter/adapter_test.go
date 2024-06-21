package adapter

import "testing"

func TestAliyunClientAdapterCreate(t *testing.T) {
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	a.CreateServer(1.0, 2.0)
}

func TestAwsClientAdapterCreate(t *testing.T) {
	var a ICreateServer = &AwsClientAdapter{
		Client: AWSClient{},
	}
	a.CreateServer(1.0, 2.0)
}
