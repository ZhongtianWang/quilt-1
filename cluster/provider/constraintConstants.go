package provider

// AwsDescriptions enumerates Amazon EC2 instance offerings
//
// T2 instances are not supported for Spot requests:
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-limits.html
var AwsDescriptions = []Description{
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "us-east-1", Price: 0.12},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "us-east-1", Price: 0.239},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "us-east-1", Price: 0.479},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "us-east-1", Price: 0.958},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "us-east-1", Price: 2.394},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "us-east-1", Price: 0.067},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "us-east-1", Price: 0.133},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "us-east-1", Price: 0.266},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "us-east-1", Price: 0.532},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "us-east-1", Price: 0.105},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "us-east-1", Price: 0.209},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "us-east-1", Price: 0.419},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "us-east-1", Price: 0.838},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "us-east-1", Price: 1.675},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "us-east-1", Price: 0.105},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "us-east-1", Price: 0.21},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "us-east-1", Price: 0.42},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "us-east-1", Price: 0.84},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "us-east-1", Price: 1.68},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "us-east-1", Price: 0.65},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "us-east-1", Price: 2.6},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "us-east-1", Price: 0.166},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "us-east-1", Price: 0.333},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "us-east-1", Price: 0.665},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "us-east-1", Price: 1.33},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "us-east-1", Price: 2.66},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "us-east-1", Price: 0.853},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "us-east-1", Price: 1.705},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "us-east-1", Price: 3.41},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "us-east-1", Price: 6.82},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "us-east-1", Price: 0.69},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "us-east-1", Price: 1.38},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "us-east-1", Price: 2.76},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "us-east-1", Price: 5.52},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "us-west-2", Price: 0.12},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "us-west-2", Price: 0.239},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "us-west-2", Price: 0.479},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "us-west-2", Price: 0.958},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "us-west-2", Price: 2.394},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "us-west-2", Price: 0.067},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "us-west-2", Price: 0.133},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "us-west-2", Price: 0.266},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "us-west-2", Price: 0.532},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "us-west-2", Price: 0.105},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "us-west-2", Price: 0.209},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "us-west-2", Price: 0.419},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "us-west-2", Price: 0.838},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "us-west-2", Price: 1.675},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "us-west-2", Price: 0.105},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "us-west-2", Price: 0.21},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "us-west-2", Price: 0.42},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "us-west-2", Price: 0.84},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "us-west-2", Price: 1.68},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "us-west-2", Price: 0.65},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "us-west-2", Price: 2.6},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "us-west-2", Price: 0.166},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "us-west-2", Price: 0.333},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "us-west-2", Price: 0.665},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "us-west-2", Price: 1.33},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "us-west-2", Price: 2.66},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "us-west-2", Price: 0.853},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "us-west-2", Price: 1.705},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "us-west-2", Price: 3.41},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "us-west-2", Price: 6.82},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "us-west-2", Price: 0.69},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "us-west-2", Price: 1.38},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "us-west-2", Price: 2.76},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "us-west-2", Price: 5.52},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "us-west-1", Price: 0.14},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "us-west-1", Price: 0.279},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "us-west-1", Price: 0.559},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "us-west-1", Price: 1.117},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "us-west-1", Price: 2.793},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "us-west-1", Price: 0.077},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "us-west-1", Price: 0.154},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "us-west-1", Price: 0.308},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "us-west-1", Price: 0.616},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "us-west-1", Price: 0.131},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "us-west-1", Price: 0.262},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "us-west-1", Price: 0.524},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "us-west-1", Price: 1.049},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "us-west-1", Price: 2.098},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "us-west-1", Price: 0.12},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "us-west-1", Price: 0.239},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "us-west-1", Price: 0.478},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "us-west-1", Price: 0.956},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "us-west-1", Price: 1.912},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "us-west-1", Price: 0.702},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "us-west-1", Price: 2.808},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "us-west-1", Price: 0.185},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "us-west-1", Price: 0.371},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "us-west-1", Price: 0.741},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "us-west-1", Price: 1.482},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "us-west-1", Price: 2.964},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "us-west-1", Price: 0.938},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "us-west-1", Price: 1.876},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "us-west-1", Price: 3.751},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "us-west-1", Price: 7.502},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "eu-west-1", Price: 0.132},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "eu-west-1", Price: 0.264},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "eu-west-1", Price: 0.528},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "eu-west-1", Price: 1.056},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "eu-west-1", Price: 2.641},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "eu-west-1", Price: 0.073},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "eu-west-1", Price: 0.146},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "eu-west-1", Price: 0.293},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "eu-west-1", Price: 0.585},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "eu-west-1", Price: 0.119},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "eu-west-1", Price: 0.238},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "eu-west-1", Price: 0.477},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "eu-west-1", Price: 0.953},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "eu-west-1", Price: 1.906},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "eu-west-1", Price: 0.12},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "eu-west-1", Price: 0.239},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "eu-west-1", Price: 0.478},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "eu-west-1", Price: 0.956},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "eu-west-1", Price: 1.912},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "eu-west-1", Price: 0.702},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "eu-west-1", Price: 2.808},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "eu-west-1", Price: 0.185},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "eu-west-1", Price: 0.371},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "eu-west-1", Price: 0.741},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "eu-west-1", Price: 1.482},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "eu-west-1", Price: 2.964},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "eu-west-1", Price: 0.938},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "eu-west-1", Price: 1.876},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "eu-west-1", Price: 3.751},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "eu-west-1", Price: 7.502},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "eu-west-1", Price: 0.735},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "eu-west-1", Price: 1.47},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "eu-west-1", Price: 2.94},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "eu-west-1", Price: 5.88},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "eu-central-1", Price: 0.143},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "eu-central-1", Price: 0.285},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "eu-central-1", Price: 0.57},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "eu-central-1", Price: 1.14},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "eu-central-1", Price: 2.85},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "eu-central-1", Price: 0.079},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "eu-central-1", Price: 0.158},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "eu-central-1", Price: 0.315},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "eu-central-1", Price: 0.632},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "eu-central-1", Price: 0.134},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "eu-central-1", Price: 0.267},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "eu-central-1", Price: 0.534},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "eu-central-1", Price: 1.069},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "eu-central-1", Price: 2.138},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "eu-central-1", Price: 0.129},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "eu-central-1", Price: 0.258},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "eu-central-1", Price: 0.516},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "eu-central-1", Price: 1.032},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "eu-central-1", Price: 2.064},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "eu-central-1", Price: 0.772},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "eu-central-1", Price: 3.088},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "eu-central-1", Price: 0.2},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "eu-central-1", Price: 0.4},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "eu-central-1", Price: 0.8},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "eu-central-1", Price: 1.6},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "eu-central-1", Price: 3.201},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "eu-central-1", Price: 1.013},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "eu-central-1", Price: 2.026},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "eu-central-1", Price: 4.051},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "eu-central-1", Price: 8.102},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "eu-central-1", Price: 0.794},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "eu-central-1", Price: 1.588},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "eu-central-1", Price: 3.176},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "eu-central-1", Price: 6.352},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "ap-southeast-1", Price: 0.178},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "ap-southeast-1", Price: 0.355},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "ap-southeast-1", Price: 0.711},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "ap-southeast-1", Price: 1.421},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "ap-southeast-1", Price: 3.553},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "ap-southeast-1", Price: 0.098},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "ap-southeast-1", Price: 0.196},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "ap-southeast-1", Price: 0.392},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "ap-southeast-1", Price: 0.784},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "ap-southeast-1", Price: 0.144},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "ap-southeast-1", Price: 0.289},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "ap-southeast-1", Price: 0.578},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "ap-southeast-1", Price: 1.155},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "ap-southeast-1", Price: 2.31},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "ap-southeast-1", Price: 0.132},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "ap-southeast-1", Price: 0.265},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "ap-southeast-1", Price: 0.529},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "ap-southeast-1", Price: 1.058},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "ap-southeast-1", Price: 2.117},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "ap-southeast-1", Price: 1},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "ap-southeast-1", Price: 4},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "ap-southeast-1", Price: 0.2},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "ap-southeast-1", Price: 0.399},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "ap-southeast-1", Price: 0.798},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "ap-southeast-1", Price: 1.596},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "ap-southeast-1", Price: 3.192},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "ap-southeast-1", Price: 1.018},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "ap-southeast-1", Price: 2.035},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "ap-southeast-1", Price: 4.07},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "ap-southeast-1", Price: 8.14},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "ap-southeast-1", Price: 0.87},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "ap-southeast-1", Price: 1.74},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "ap-southeast-1", Price: 3.48},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "ap-southeast-1", Price: 6.96},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "ap-northeast-1", Price: 0.174},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "ap-northeast-1", Price: 0.348},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "ap-northeast-1", Price: 0.695},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "ap-northeast-1", Price: 1.391},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "ap-northeast-1", Price: 3.477},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "ap-northeast-1", Price: 0.096},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "ap-northeast-1", Price: 0.193},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "ap-northeast-1", Price: 0.385},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "ap-northeast-1", Price: 0.77},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "ap-northeast-1", Price: 0.133},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "ap-northeast-1", Price: 0.265},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "ap-northeast-1", Price: 0.531},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "ap-northeast-1", Price: 1.061},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "ap-northeast-1", Price: 2.122},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "ap-northeast-1", Price: 0.128},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "ap-northeast-1", Price: 0.255},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "ap-northeast-1", Price: 0.511},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "ap-northeast-1", Price: 1.021},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "ap-northeast-1", Price: 2.043},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "ap-northeast-1", Price: 0.898},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "ap-northeast-1", Price: 3.592},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "ap-northeast-1", Price: 0.2},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "ap-northeast-1", Price: 0.399},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "ap-northeast-1", Price: 0.798},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "ap-northeast-1", Price: 1.596},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "ap-northeast-1", Price: 3.192},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "ap-northeast-1", Price: 1.001},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "ap-northeast-1", Price: 2.001},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "ap-northeast-1", Price: 4.002},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "ap-northeast-1", Price: 8.004},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "ap-northeast-1", Price: 0.844},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "ap-northeast-1", Price: 1.688},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "ap-northeast-1", Price: 3.376},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "ap-northeast-1", Price: 6.752},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "ap-southeast-2", Price: 0.168},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "ap-southeast-2", Price: 0.336},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "ap-southeast-2", Price: 0.673},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "ap-southeast-2", Price: 1.345},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "ap-southeast-2", Price: 3.363},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "ap-southeast-2", Price: 0.093},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "ap-southeast-2", Price: 0.186},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "ap-southeast-2", Price: 0.372},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "ap-southeast-2", Price: 0.745},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "ap-southeast-2", Price: 0.137},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "ap-southeast-2", Price: 0.275},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "ap-southeast-2", Price: 0.549},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "ap-southeast-2", Price: 1.097},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "ap-southeast-2", Price: 2.195},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "ap-southeast-2", Price: 0.132},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "ap-southeast-2", Price: 0.265},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "ap-southeast-2", Price: 0.529},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "ap-southeast-2", Price: 1.058},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "ap-southeast-2", Price: 2.117},
	{Size: "g2.2xlarge", CPU: 8, RAM: 15, Disk: "60 SSD", Region: "ap-southeast-2", Price: 0.898},
	{Size: "g2.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 120 SSD", Region: "ap-southeast-2", Price: 3.592},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "ap-southeast-2", Price: 0.2},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "ap-southeast-2", Price: 0.399},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "ap-southeast-2", Price: 0.798},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "ap-southeast-2", Price: 1.596},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "ap-southeast-2", Price: 3.192},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "ap-southeast-2", Price: 1.018},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "ap-southeast-2", Price: 2.035},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "ap-southeast-2", Price: 4.07},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "ap-southeast-2", Price: 8.14},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "ap-southeast-2", Price: 0.87},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "ap-southeast-2", Price: 1.74},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "ap-southeast-2", Price: 3.48},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "ap-southeast-2", Price: 6.96},
	{Size: "m4.large", CPU: 2, RAM: 8, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.165},
	{Size: "m4.xlarge", CPU: 4, RAM: 16, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.331},
	{Size: "m4.2xlarge", CPU: 8, RAM: 32, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.660},
	{Size: "m4.4xlarge", CPU: 16, RAM: 64, Disk: "ebsonly", Region: "ap-northeast-2", Price: 1.321},
	{Size: "m4.10xlarge", CPU: 40, RAM: 160, Disk: "ebsonly", Region: "ap-northeast-2", Price: 3.303},
	{Size: "c4.large", CPU: 2, RAM: 3.75, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.120},
	{Size: "c4.xlarge", CPU: 4, RAM: 7.5, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.239},
	{Size: "c4.2xlarge", CPU: 8, RAM: 15, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.478},
	{Size: "c4.4xlarge", CPU: 16, RAM: 30, Disk: "ebsonly", Region: "ap-northeast-2", Price: 0.955},
	{Size: "c4.8xlarge", CPU: 36, RAM: 60, Disk: "ebsonly", Region: "ap-northeast-2", Price: 1.910},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "ap-northeast-2", Price: 0.200},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "ap-northeast-2", Price: 0.399},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "ap-northeast-2", Price: 0.798},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "ap-northeast-2", Price: 1.596},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "ap-northeast-2", Price: 3.192},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "ap-northeast-2", Price: 1.001},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "ap-northeast-2", Price: 2.001},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "ap-northeast-2", Price: 4.002},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "ap-northeast-2", Price: 8.004},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "ap-northeast-2", Price: 0.844},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "ap-northeast-2", Price: 1.688},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "ap-northeast-2", Price: 3.376},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "ap-northeast-2", Price: 6.752},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "sa-east-1", Price: 0.095},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "sa-east-1", Price: 0.19},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "sa-east-1", Price: 0.381},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "sa-east-1", Price: 0.761},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "sa-east-1", Price: 0.163},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "sa-east-1", Price: 0.325},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "sa-east-1", Price: 0.65},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "sa-east-1", Price: 1.3},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "sa-east-1", Price: 2.6},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "sa-east-1", Price: 2.799},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "sa-east-1", Price: 5.597},
	{Size: "m3.medium", CPU: 1, RAM: 3.75, Disk: "1 x 4 SSD", Region: "us-gov-west-1", Price: 0.084},
	{Size: "m3.large", CPU: 2, RAM: 7.5, Disk: "1 x 32 SSD", Region: "us-gov-west-1", Price: 0.168},
	{Size: "m3.xlarge", CPU: 4, RAM: 15, Disk: "2 x 40 SSD", Region: "us-gov-west-1", Price: 0.336},
	{Size: "m3.2xlarge", CPU: 8, RAM: 30, Disk: "2 x 80 SSD", Region: "us-gov-west-1", Price: 0.672},
	{Size: "c3.large", CPU: 2, RAM: 3.75, Disk: "2 x 16 SSD", Region: "us-gov-west-1", Price: 0.126},
	{Size: "c3.xlarge", CPU: 4, RAM: 7.5, Disk: "2 x 40 SSD", Region: "us-gov-west-1", Price: 0.252},
	{Size: "c3.2xlarge", CPU: 8, RAM: 15, Disk: "2 x 80 SSD", Region: "us-gov-west-1", Price: 0.504},
	{Size: "c3.4xlarge", CPU: 16, RAM: 30, Disk: "2 x 160 SSD", Region: "us-gov-west-1", Price: 1.008},
	{Size: "c3.8xlarge", CPU: 32, RAM: 60, Disk: "2 x 320 SSD", Region: "us-gov-west-1", Price: 2.016},
	{Size: "r3.large", CPU: 2, RAM: 15, Disk: "1 x 32 SSD", Region: "us-gov-west-1", Price: 0.2},
	{Size: "r3.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 80 SSD", Region: "us-gov-west-1", Price: 0.399},
	{Size: "r3.2xlarge", CPU: 8, RAM: 61, Disk: "1 x 160 SSD", Region: "us-gov-west-1", Price: 0.798},
	{Size: "r3.4xlarge", CPU: 16, RAM: 122, Disk: "1 x 320 SSD", Region: "us-gov-west-1", Price: 1.596},
	{Size: "r3.8xlarge", CPU: 32, RAM: 244, Disk: "2 x 320 SSD", Region: "us-gov-west-1", Price: 3.192},
	{Size: "i2.xlarge", CPU: 4, RAM: 30.5, Disk: "1 x 800 SSD", Region: "us-gov-west-1", Price: 1.023},
	{Size: "i2.2xlarge", CPU: 8, RAM: 61, Disk: "2 x 800 SSD", Region: "us-gov-west-1", Price: 2.046},
	{Size: "i2.4xlarge", CPU: 16, RAM: 122, Disk: "4 x 800 SSD", Region: "us-gov-west-1", Price: 4.092},
	{Size: "i2.8xlarge", CPU: 32, RAM: 244, Disk: "8 x 800 SSD", Region: "us-gov-west-1", Price: 8.184},
	{Size: "d2.xlarge", CPU: 4, RAM: 30.5, Disk: "3 x 2000 HDD", Region: "us-gov-west-1", Price: 0.828},
	{Size: "d2.2xlarge", CPU: 8, RAM: 61, Disk: "6 x 2000 HDD", Region: "us-gov-west-1", Price: 1.656},
	{Size: "d2.4xlarge", CPU: 16, RAM: 122, Disk: "12 x 2000 HDD", Region: "us-gov-west-1", Price: 3.312},
	{Size: "d2.8xlarge", CPU: 36, RAM: 244, Disk: "24 x 2000 HDD", Region: "us-gov-west-1", Price: 6.624},
}

// AzureDescriptions enumerates Microsoft Azure instance offerings TODO
var AzureDescriptions = []Description{}
