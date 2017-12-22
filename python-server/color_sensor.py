import time
import smbus
import Adafruit_TCS34725

def sense_colors():
	# Create a TCS34725 instance with default integration time (2.4ms) and gain (4x).
	tcs = Adafruit_TCS34725.TCS34725()

	# Disable interrupts (can enable them by passing true, see the set_interrupt_limits function too).
	tcs.set_interrupt(False)

	# Read the R, G, B, C color data.
	r, g, b, c = tcs.get_raw_data()

	# Enable interrupts and put the chip back to low power sleep/disabled.
	tcs.set_interrupt(True)
	tcs.disable()

	return r, g, b
