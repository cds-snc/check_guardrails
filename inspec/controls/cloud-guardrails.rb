# copyright: 2020, Max Neuvians

title "PBMM Cloud Guardrails"

control "PBMM-cloud-guardrails-1-0" do
  impact 0.7                                
  title "Protect Root / Global Admins Account"             
  desc "Validates that there is no access key for the AWS account’s root user"
  describe aws_iam_root_user do
    it { should_not have_access_key }
  end
end

control "PBMM-cloud-guardrails-1-1" do
  impact 0.7                                
  title "Protect Root / Global Admins Account"             
  desc "Verify breakglass accounts."
  describe aws_iam_users
    .where( has_mfa_enabled: false)
    .where( has_console_password: true ) do
    its('entries.count') { should eq 2 }
  end
end

control "PBMM-cloud-guardrails-1-3" do
  impact 0.7                                
  title "Protect Root / Global Admins Account"             
  desc "Ensure MFA is enabled for the root account"
  describe aws_iam_root_user do
    it { should have_mfa_enabled }
  end
end

control "PBMM-cloud-guardrails-2-4" do
  impact 0.7                                
  title "Management of administrative privileges"             
  desc "Verify a strong password policy is enabled."
  describe aws_iam_password_policy do
    its('minimum_password_length') { should be >= 14 }
  end
end

control "PBMM-cloud-guardrails-5-0" do
  impact 0.7                                
  title "Cloud Console Access (Developers/ApplicationOwners)"             
  desc "Ensure MFA is enabled for all IAM accounts with passwords (except breakglass)"
  describe aws_iam_users
    .where( has_mfa_enabled: false )
    .where( has_console_password: true ) do
    its('entries.count') { should eq 2 }
  end
end

control "PBMM-cloud-guardrails-5-1" do
  impact 0.7                                
  title "Cloud Console Access (Developers/ApplicationOwners)"             
  desc "Ensure no non-MFA account has passwords (except breakglass)"
  describe aws_iam_users
    .where( has_mfa_enabled: false )
    .where( has_console_password: true ) do
    its('entries.count') { should eq 2 }
  end
end

control "PBMM-cloud-guardrails-7-0-ec2" do
  impact 0.7                                
  title "Data location in Canada"             
  desc "Verify that all EC2 instances are deployed only within the AWS Canada Central Region"
  aws_regions.region_names.each do |region| 
    describe aws_ec2_instances(aws_region: region) do
      if region != "ca-central-1" 
        its('entries.count') { should cmp 0 }
      end
    end
  end
end

control "PBMM-cloud-guardrails-7-0-rds" do
  impact 0.7                                
  title "Data location in Canada"             
  desc "Verify that all RDS instances are deployed only within the AWS Canada Central Region"
  aws_regions.region_names.each do |region| 
    describe aws_rds_instances(aws_region: region) do
      if region != "ca-central-1" 
        its('entries.count') { should cmp 0 }
      end
    end
  end
end

control "PBMM-cloud-guardrails-8-0-ec2" do
  impact 0.7                                
  title "Protection of data-at-rest"             
  desc "Verify that all EC2 disks are encrypted"
  aws_regions.region_names.each do |region| 
    aws_ebs_volumes(aws_region: region).volume_ids.each do |volume_id|
      describe aws_ebs_volume(volume_id) do
        it { should be_encrypted }
      end
    end
  end
end

control "PBMM-cloud-guardrails-8-0-rds" do
  impact 0.7                                
  title "Protection of data-at-rest"             
  desc "Verify that all RDS disks are encrypted"
  aws_regions.region_names.each do |region| 
    aws_rds_instances(aws_region: region).db_instance_identifiers.each do |db_instance_identifier|
      describe aws_rds_instance(db_instance_identifier) do
        it { should be_encrypted }
      end
    end
  end
end

control "PBMM-cloud-guardrails-8-0-s3" do
  impact 0.7                                
  title "Protection of data-at-rest"             
  desc "Verify that all S3 buckets are encrypted"
  aws_s3_buckets.bucket_names.each do |bucket_name|
    describe aws_s3_bucket(bucket_name) do
      it { should have_default_encryption_enabled }
    end
  end
end

control "PBMM-cloud-guardrails-9-0" do
  impact 0.7                                
  title "Protection of data-in-transit"             
  desc "Verify that no security groups allow port 80 ingress from the internet"
  aws_regions.region_names.each do |region|
    aws_security_groups(aws_region: region).group_ids.each do |group_id|
      describe aws_security_group(group_id) do
        it { should_not allow_in(port: 80, ipv4_range: '0.0.0.0/0') }
      end
    end
  end
end

control "PBMM-cloud-guardrails-11-0" do
  impact 0.7                                
  title "Logging and Montitoring"             
  desc "Has a AWS-Landing-Zone-BaselineCloudTrail cloud trail"
  describe aws_cloudtrail_trails do
    its('names') { should include /AWS\-Landing\-Zone\-BaselineCloudTrail/ }
  end
end

control "PBMM-cloud-guardrails-11-1" do
  impact 0.7                                
  title "Logging and Montitoring"             
  desc "Has a AWS-Landing-Zone-Security-Notification SNS topic"
  describe aws_sns_topics do
    its('topic_arns') { should include /AWS\-Landing\-Zone\-Security\-Notification/ }
  end
end
