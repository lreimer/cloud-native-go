#!/bin/sh
cat > $APP_NAME.json <<EOF
[{
    "id": "$APP_NAME",
    "instances": 3,
    "container": {
        "type": "DOCKER",
        "docker": {
            "image": "lreimer/cloud-native-go:$WERCKER_GIT_COMMIT",
            "network": "BRIDGE",
            "forcePullImage": true,
            "privileged": false,
            "portMappings": [
                {
                    "containerPort": 8080,
                    "protocol": "tcp",
                    "labels": {
                        "VIP_0": "cloud-native-go:8080"
                    }
                }
            ]
        }
    },
    "healthChecks": [{
        "protocol": "TCP",
        "gracePeriodSeconds": 600,
        "intervalSeconds": 30,
        "portIndex": 0,
        "timeoutSeconds": 10,
        "maxConsecutiveFailures": 2
    }],
    "cpus": 0.5,
    "mem": 64.0,
    "disk": 0
}]
EOF