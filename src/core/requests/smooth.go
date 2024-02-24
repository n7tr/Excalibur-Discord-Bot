package requests

func Smooth[T any](channels []T) [][]T {
	const step = 50
	smoothed := make([][]T, 0, len(channels)/step+1)
	for i := 0; i < len(channels); i += step {
		end := i + step
		if end > len(channels) {
			end = len(channels)
		}
		smoothed = append(smoothed, channels[i:end])
	}
	return smoothed
}
