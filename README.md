# FFmpeg Scripts

These scripts use FFmpeg to transcode video files with specified parameters.

## Python Script

### Requirements

- Python 3.x
- FFmpeg installed and added to system PATH

### Installation

No installation required beyond Python and FFmpeg setup.

### Usage

1. Modify the `run_ffmpeg` function in `ffmpeg.py` to adjust encoding parameters.
2. Run the script using Python:

    ```bash
    python ffmpeg.py
    ```

3. Adjust input and output file paths as needed within the script.

### Parameters

- `input_file`: Path to the input video file.
- `output_file`: Path to the output video file.

### Notes

- Modify the FFmpeg command (`command` list) within the script to change encoding options such as video codec, audio codec, bitrate, resolution, etc.
- Uncomment or modify parameters in the `command` list to apply specific filters, change encoding presets, quality scale, or map specific streams from the input file.

## Go Script

### Requirements

- Go programming language installed.
- FFmpeg installed and added to system PATH.

### Installation

No additional installation required beyond Go and FFmpeg setup.

### Usage

1. Modify the `runFFmpeg` function in `main.go` to adjust encoding parameters.
2. Run the script using Go:

    ```bash
    go run ffmpeg.go
    ```

3. Adjust input and output file paths as needed within the script.

### Parameters

- `inputFile`: Path to the input video file.
- `outputFile`: Path to the output video file.

### Notes

- Modify the `command` slice within the `runFFmpeg` function to change encoding options such as video codec, audio codec, bitrate, resolution, etc.
- Uncomment or modify parameters in the `command` slice to apply specific filters, change encoding presets, quality scale, or map specific streams from the input file.


