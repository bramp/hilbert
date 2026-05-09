# Space-Filling Curves Implementation Roadmap

This document tracks the progress of adding new space-filling curves to the `hilbert` package. Each implementation must include:
- [ ] Core implementation in a new file (e.g., `morton.go`).
- [ ] `Map` and `MapInverse` functions.
- [ ] Comprehensive tests (including round-trip tests).
- [ ] A sample image generated and stored in `images/`.
- [ ] Updated `README.md` with descriptions and visuals.

## Curves to Implement

### 1. Morton Order (Z-Order)
- [x] Implementation (`morton.go`)
- [x] Tests (`morton_test.go`)
- [x] Sample image (`images/morton.png`)
- [x] README update

### 2. Moore Curve
- [x] Implementation (`moore.go`)
- [x] Tests (`moore_test.go`)
- [x] Sample image (`images/moore.png`)
- [x] README update

### 3. Gosper Curve
- [ ] Implementation (`gosper.go`)
- [ ] Tests (`gosper_test.go`)
- [ ] Sample image (`images/gosper.png`)
- [ ] README update

### 4. Sierpinski Curve
- [ ] Implementation (`sierpinski.go`)
- [ ] Tests (`sierpinski_test.go`)
- [ ] Sample image (`images/sierpinski.png`)
- [ ] README update

### 5. Dragon Curve
- [ ] Implementation (`dragon.go`)
- [ ] Tests (`dragon_test.go`)
- [ ] Sample image (`images/dragon.png`)
- [ ] README update

## Documentation & Cleanup
- [ ] Reorganize `README.md` to use a gallery/table for curves.
- [ ] Ensure all curves implement the `SpaceFilling` interface.
- [ ] Add a common interface definition if not already explicit.
