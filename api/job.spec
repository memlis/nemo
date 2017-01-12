{
  "name": "testjob",
  "type": "service",
  "user": "testuser",
  "replicas": 10,
  "image": "nginx:1.10",
  "cpu": 0.1,
  "gpu": 0,
  "memory": 8,
  "disk": 0,
  "network": "BRIDGE",
  "forcePullImage": true,
  "privileged": true,
  "priority": 100,
  "environment": {
    "key": "value", 
  },
  "command":"",
  "labels": {
    "key": "value",
  },
  "ports": [
    {
        "port": 80,
        "protocol": "tcp",
        "name": web
    }
  ],
  "volumes": [
    {   
      "hostPath": "/home",
      "containerPath": "/data",
      "mode": "RW"
    }   
  ],
  "uris": [
    {
      "uri": "https://nginx.org/download/nginx-1.8.1.tar.gz",
      "executable": false,
      "extract": false,
      "cache": false,
    }
  ],
  "constraints": [
    {
      "attribute": "cluster",
      "value": "dataman",
    }
  ],
  "restart": {
    "attempts": 3,
    "delay": 10
  },
  "update": {
  },
  "kill": {
    duration: 5
  }
}

{
    "id": "taurusjob2",
    "type": "servive",
    "user": "test",
    "cluster": "MyCluster",
    "tasks": [
        {   
            "environment": "prod",
            "role": "*",
            "container": {
                "image": "nginx:1.10"
            },  
            "resources": {
                "cpu": 0.1,
                "memory": 10  
            },  
            "replicas": 4
        }   
    ]   
}

