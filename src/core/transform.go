package core

import (
	"math"
)

// Transform is a 4x4 transformation matrix.
type Transform struct {
	mat [4][4]Float // Matrix entries
}

// NewTransform returns a pointer of the Transform.
func NewTransform(mat [4][4]Float) *Transform {
	t := new(Transform)
	t.mat = mat
	return t
}

// NewTransformWithElements returns a pointer of the Transform,
// which has specified elements as its entries.
func NewTransformWithElements(
	m00, m01, m02, m03 Float,
	m10, m11, m12, m13 Float,
	m20, m21, m22, m23 Float,
	m30, m31, m32, m33 Float) *Transform {
	t := new(Transform)
	t.mat = [4][4]Float{
		{m00, m01, m02, m03},
		{m10, m11, m12, m13},
		{m20, m21, m22, m23},
		{m30, m31, m32, m33},
	}
	return t
}

// NewScale returns a scaling Transform.
func NewScale(sx, sy, sz Float) *Transform {
	return NewTransformWithElements(
		sx, 0.0, 0.0, 0.0,
		0.0, sy, 0.0, 0.0,
		0.0, 0.0, sz, 0.0,
		0.0, 0.0, 0.0, 1.0)
}

// NewLookAt returns a pointer of the "lookAt" Transform.
func NewLookAt(origin, target, up *Vector3d) *Transform {
	m03 := origin.X
	m13 := origin.Y
	m23 := origin.Z
	m33 := 1.0

	dir := target.Subtract(origin).Normalized()
	left := dir.Cross(up).Normalized()
	newUp := left.Cross(dir)

	m00 := left.X
	m10 := left.Y
	m20 := left.Z
	m30 := 0.0
	m01 := newUp.X
	m11 := newUp.Y
	m21 := newUp.Z
	m31 := 0.0
	m02 := dir.X
	m12 := dir.Y
	m22 := dir.Z
	m32 := 0.0

	return NewTransformWithElements(
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33,
	)
}

// NewPerspective returns a pointer of the perspective Transform.
func NewPerspective(fov, aspect, near, far Float) *Transform {
	pers := [4][4]Float{
		{1.0 / aspect, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, far / (far - near), -far * near / (far - near)},
		{0.0, 0.0, 1.0, 0.0},
	}
	s := 1.0 / math.Tan(DegreeToRadian(fov)*0.5)
	return NewScale(s, s, 1.0).Multiply(NewTransform(pers))
}

// At returns an element of Transform.
func (t *Transform) At(i, j int) Float {
	if i < 0 || j < 0 || i >= 4 || j >= 4 {
		panic("Index out of bounds!")
	}
	return t.mat[i][j]
}

// Multiply computes the multiplications of Transforms.
func (t1 *Transform) Multiply(t2 *Transform) *Transform {
	ret := new(Transform)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret.mat[i][j] = 0.0
			for k := 0; k < 4; k++ {
				ret.mat[i][j] += t1.mat[i][k] * t2.mat[k][j]
			}
		}
	}
	return ret
}

// Apply multiplies the transform to a vector.
func (t *Transform) ApplyToP(v *Vector3d) *Vector3d {
	u := [4]Float{v.X, v.Y, v.Z, 1.0}
	ret := [4]Float{0.0, 0.0, 0.0, 0.0}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret[i] += t.mat[i][j] * u[j]
		}
	}
	return NewVector3d(ret[0], ret[1], ret[2]).Divide(ret[3])
}

// Apply multiplies the transform to a vector.
func (t *Transform) ApplyToV(v *Vector3d) *Vector3d {
	u := [3]Float{v.X, v.Y, v.Z}
	ret := [3]Float{0.0, 0.0, 0.0}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret[i] += t.mat[i][j] * u[j]
		}
	}
	return NewVector3d(ret[0], ret[1], ret[2])
}

// Inverted returns a inverted Transform.
func (t *Transform) Inverted() *Transform {
	indxc := make([]int, 4)
	indxr := make([]int, 4)
	ipiv := [4]int{0, 0, 0, 0}

	mInv := [4][4]Float{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mInv[i][j] = t.mat[i][j]
		}
	}

	for i := 0; i < 4; i++ {
		irow, icol := 0, 0
		big := 0.0

		// Choose pivot
		for j := 0; j < 4; j++ {
			if ipiv[j] != 1 {
				for k := 0; k < 4; k++ {
					if ipiv[k] == 0 {
						if math.Abs(mInv[j][k]) >= big {
							big = math.Abs(mInv[j][k])
							irow = j
							icol = k
						}
					} else if ipiv[k] > 1 {
						panic("Singular matrix cannot be inverted")
					}
				}
			}
		}
		ipiv[icol]++

		// Swap pivot row
		if irow != icol {
			for k := 0; k < 4; k++ {
				mInv[irow][k], mInv[icol][k] = mInv[icol][k], mInv[irow][k]
			}
		}

		indxr[i] = irow
		indxc[i] = icol
		if math.Abs(mInv[icol][icol]) < Eps {
			panic("Singular matrix cannot be inverted")
		}

		pinv := 1.0 / mInv[icol][icol]
		mInv[icol][icol] = 1.0
		for j := 0; j < 4; j++ {
			mInv[icol][j] *= pinv
		}

		// Subtract diagonal value from the other rows
		for j := 0; j < 4; j++ {
			if j != icol {
				save := mInv[j][icol]
				mInv[j][icol] = 0.0
				for k := 0; k < 4; k++ {
					mInv[j][k] -= mInv[icol][k] * save
				}
			}
		}
	}

	// Swap columins to reset permutations
	for j := 3; j >= 0; j-- {
		if indxr[j] != indxc[j] {
			for k := 0; k < 4; k++ {
				mInv[k][indxr[j]], mInv[k][indxc[j]] =
					mInv[k][indxc[j]], mInv[k][indxr[j]]
			}
		}
	}

	return NewTransform(mInv)
}
