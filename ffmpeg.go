package main

import (
	"log"
	"os/exec"
)

func runFFmpeg(inputFile, outputFile string) {
	command := []string{
		"ffmpeg",

		// Input file
		"-i", inputFile, // Specify the input file

		// Video codec
		"-c:v", "libx264", // Video codec to use for encoding (H.264)
		// "-c:v", "copy", // Uncomment to copy the video stream without re-encoding

		// Audio codec
		"-c:a", "aac", // Audio codec to use for encoding (AAC)
		// "-c:a", "copy", // Uncomment to copy the audio stream without re-encoding

		// Bitrate
		"-b:v", "1000k", // Video bitrate (in kbps)
		"-b:a", "128k", // Audio bitrate (in kbps)

		// Resolution
		"-s", "1280x720", // Set the video resolution (width x height)

		// Frame rate
		"-r", "30", // Set the frame rate (frames per second)

		// Aspect ratio
		"-aspect", "16:9", // Set the aspect ratio

		// Audio sample rate
		"-ar", "44100", // Set the audio sample rate (in Hz)

		// Channels
		"-ac", "2", // Set the number of audio channels

		// Audio volume
		"-af", "volume=1.0", // Set audio volume (2.0 means double the volume)

		// Video filters
		"-vf", "scale=1280:720", // Apply video filter (e.g., scaling)
		// "-vf", "crop=640:360:320:180", // Uncomment to crop the video (width:height:x:y)

		// Additional options

		// Bitstream filters
		// "-bsf:v", "h264_mp4toannexb", // Apply bitstream filter to video stream

		// Codec-specific options
		// "-preset", "fast", // Set the encoding preset for libx264 (e.g., ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow, placebo)
		// "-crf", "23", // Constant Rate Factor (0-51, lower means better quality)

		// Quality scale
		// "-q:v", "2", // Set the video quality scale (1-31, lower means better quality)

		// Audio filters
		// "-af", "highpass=f=200", // Apply a high-pass filter to the audio (frequency in Hz)
		// "-af", "lowpass=f=3000", // Apply a low-pass filter to the audio (frequency in Hz)

		// Subtitle options
		// "-scodec", "mov_text", // Specify the subtitle codec
		// "-sn", // Disable subtitle stream

		// Map streams
		// "-map", "0:v:0", // Map the first video stream from the input
		// "-map", "0:a:0", // Map the first audio stream from the input

		// Start time
		// "-ss", "00:00:30", // Start the output from the 30th second of the input

		// Duration
		// "-t", "00:01:00", // Limit the output to the first 1 minute of the input

		// Overwrite output file if it exists
		"-y",

		// Output file
		outputFile,
	}

	cmd := exec.Command(command[0], command[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runFFmpeg("input.mp4", "output.mp4")
}
