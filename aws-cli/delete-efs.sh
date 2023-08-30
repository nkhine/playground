for fs in $(aws efs describe-file-systems --profile ab-dev --region eu-west-1 --query 'FileSystems[*].FileSystemId' --output text); do
    for mt in $(aws efs describe-mount-targets --profile ab-dev --region eu-west-1 --file-system-id $fs --query 'MountTargets[*].MountTargetId' --output text); do
        aws efs delete-mount-target --profile ab-dev --region eu-west-1  --mount-target-id $mt
    done
    aws efs delete-file-system --profile ab-dev --region eu-west-1  --file-system-id $fs
done
