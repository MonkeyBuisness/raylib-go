// Useful easing functions for values animation
//
// A port of Robert Penner's easing equations (http://robertpenner.com/easing/)
package easings

import (
	"math"
)

// Linear Easing functions

// Linear None
func LinearNone(t, b, c, d float32) float32 {
	return c*t/d + b
}

// Linear In
func LinearIn(t, b, c, d float32) float32 {
	return c*t/d + b
}

// Linear Out
func LinearOut(t, b, c, d float32) float32 {
	return c*t/d + b
}

// Linear In Out
func LinearInOut(t, b, c, d float32) float32 {
	return c*t/d + b
}

// Sine Easing functions

// Sine In
func SineIn(t, b, c, d float32) float32 {
	return -c*float32(math.Cos(float64(t/d)*(math.Pi/2))) + c + b
}

// Sine Out
func SineOut(t, b, c, d float32) float32 {
	return c*float32(math.Sin(float64(t/d)*(math.Pi/2))) + b
}

// Sine In Out
func SineInOut(t, b, c, d float32) float32 {
	return -c/2*(float32(math.Cos(math.Pi*float64(t/d)))-1) + b
}

// Circular Easing functions

// Circ In
func CircIn(t, b, c, d float32) float32 {
	t = t / d
	return -c*(float32(math.Sqrt(float64(1-t*t)))-1) + b
}

// Circ Out
func CircOut(t, b, c, d float32) float32 {
	return c*float32(math.Sqrt(1-float64((t/d-1)*t))) + b
}

// Circ In Out
func CircInOut(t, b, c, d float32) float32 {
	t = t / d * 2

	if t < 1 {
		return -c/2*(float32(math.Sqrt(float64(1-t*t)))-1) + b
	}

	t = t - 2
	return c/2*(float32(math.Sqrt(1-float64(t*t)))+1) + b
}

// Cubic Easing functions

// Cubic In
func CubicIn(t, b, c, d float32) float32 {
	t = t / d
	return c*t*t*t + b
}

// Cubic Out
func CubicOut(t, b, c, d float32) float32 {
	t = t/d - 1
	return c*(t*t*t+1) + b
}

// Cubic In Out
func CubicInOut(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return (c/2*t*t*t + b)
	}

	t = t - 2
	return c/2*(t*t*t+2) + b
}

// Quadratic Easing functions

// Quad In
func QuadIn(t, b, c, d float32) float32 {
	t = t / d
	return c*t*t + b
}

// Quad Out
func QuadOut(t, b, c, d float32) float32 {
	t = t / d
	return (-c*t*(t-2) + b)
}

// Quad In Out
func QuadInOut(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return ((c / 2) * (t * t)) + b
	}

	return -c/2*((t-1)*(t-3)-1) + b
}

// Exponential Easing functions

// Expo In
func ExpoIn(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}

	return (c*float32(math.Pow(2, 10*float64(t/d-1))) + b)
}

// Expo Out
func ExpoOut(t, b, c, d float32) float32 {
	if t == d {
		return (b + c)
	}

	return c*(-float32(math.Pow(2, -10*float64(t/d)))+1) + b
}

// Expo In Out
func ExpoInOut(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}
	if t == d {
		return (b + c)
	}

	t = t / d * 2

	if t < 1 {
		return (c/2*float32(math.Pow(2, 10*float64(t-1))) + b)
	}

	t = t - 1
	return (c/2*(-float32(math.Pow(2, -10*float64(t)))+2) + b)
}

// Back Easing functions

// Back In
func BackIn(t, b, c, d float32) float32 {
	s := float32(1.70158)
	t = t / d
	return c*t*t*((s+1)*t-s) + b
}

// Back Out
func BackOut(t, b, c, d float32) float32 {
	s := float32(1.70158)
	t = t/d - 1
	return c*(t*t*((s+1)*t+s)+1) + b
}

// Back In Out
func BackInOut(t, b, c, d float32) float32 {
	s := float32(1.70158)
	s = s * 1.525
	t = t / d * 2

	if t < 1 {
		return c/2*(t*t*((s+1)*t-s)) + b
	}

	t = t - 2
	return c/2*(t*t*((s+1)*t+s)+2) + b
}

// Bounce Easing functions

// Bounce In
func BounceIn(t, b, c, d float32) float32 {
	return (c - BounceOut(d-t, 0, c, d) + b)
}

// Bounce Out
func BounceOut(t, b, c, d float32) float32 {
	t = t / d
	if t < (1 / 2.75) {
		return (c*(7.5625*t*t) + b)
	} else if t < (2 / 2.75) {
		t = t - (1.5 / 2.75)
		return c*(7.5625*t*t+0.75) + b
	} else if t < (2.5 / 2.75) {
		t = t - (2.25 / 2.75)
		return c*(7.5625*t*t+0.9375) + b
	}

	t = t - (2.625 / 2.75)
	return c*(7.5625*t*t+0.984375) + b
}

// Bounce In Out
func BounceInOut(t, b, c, d float32) float32 {
	if t < d/2 {
		return BounceIn(t*2, 0, c, d)*0.5 + b
	}

	return BounceOut(t*2-d, 0, c, d)*0.5 + c*0.5 + b
}

// Elastic Easing functions

// Elastic In
func ElasticIn(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}

	t = t / d

	if t == 1 {
		return b + c
	}

	p := d * 0.3
	a := c
	s := p / 4
	postFix := a * float32(math.Pow(2, 10*float64(t-1)))

	return -(postFix * float32(math.Sin(float64(t*d-s)*(2*math.Pi)/float64(p)))) + b
}

// Elastic Out
func ElasticOut(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}

	t = t / d

	if t == 1 {
		return b + c
	}

	p := d * 0.3
	a := c
	s := p / 4

	return a*float32(math.Pow(2, -10*float64(t)))*float32(math.Sin(float64(t*d-s)*(2*math.Pi)/float64(p))) + c + b
}

// Elastic In Out
func ElasticInOut(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}

	t = t / d * 2

	if t == 2 {
		return b + c
	}

	p := d * (0.3 * 1.5)
	a := c
	s := p / 4

	if t < 1 {
		t = t - 1
		postFix := a * float32(math.Pow(2, 10*float64(t)))
		return -0.5*(postFix*float32(math.Sin(float64(t*d-s)*(2*math.Pi)/float64(p)))) + b
	}

	t = t - 1
	postFix := a * float32(math.Pow(2, -10*(float64(t))))
	return postFix*float32(math.Sin(float64(t*d-s)*(2*math.Pi)/float64(p)))*0.5 + c + b
}
