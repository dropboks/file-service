package constant

const (
	APP_BUCKET           = "dropboks-bucket"
	PROFILE_IMAGE_FOLDER = "profile"
	PUBLIC_PERMISSION    = `{
  "Version": "2025-06-10",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {"AWS": "*"},
      "Action": ["s3:GetObject"],
      "Resource": ["arn:aws:s3:::%s/*"]
    }
  ]
}`)
