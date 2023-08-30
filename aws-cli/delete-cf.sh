for region in $(aws ec2 describe-regions --query 'Regions[*].RegionName' --output text --profile ab-cicd); do
    echo "Processing region: $region"
    
    # List all CloudFormation stacks in the current region
    for stack in $(aws cloudformation describe-stacks --region $region --query 'Stacks[*].StackName' --output text --profile ab-cicd); do
        echo "Deleting stack: $stack in region: $region"
        
        # Delete the CloudFormation stack
        aws cloudformation delete-stack --stack-name $stack --region $region --profile ab-cicd
        
        # Optional: Wait for the stack to be deleted before moving to the next one
        aws cloudformation wait stack-delete-complete --stack-name $stack --region $region --profile ab-cicd
    done
done
