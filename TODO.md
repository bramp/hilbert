# Space-Filling Curves Implementation Roadmap

This document tracks the progress of adding new space-filling curves to the `hilbert` package.

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

### 3. Snake Traversal (Boustrophedon)
- [ ] Implementation (`snake.go`)
- [ ] Tests (`snake_test.go`)
- [ ] Sample image (`images/snake.png`)
- [ ] README update

### 4. Gosper Curve (Hexagonal)
- [ ] Implementation (`gosper.go`)
- [ ] Tests (`gosper_test.go`)
- [ ] Sample image (`images/gosper.png`)
- [ ] README update

### 5. Sierpinski Arrowhead (Triangular)
- [ ] Implementation (`sierpinski.go`)
- [ ] Tests (`sierpinski_test.go`)
- [ ] Sample image (`images/sierpinski.png`)
- [ ] README update

### 6. Dragon Curve (Fractal)
- [ ] Implementation (`dragon.go`)
- [ ] Tests (`dragon_test.go`)
- [ ] Sample image (`images/dragon.png`)
- [ ] README update

## Documentation & Cleanup
- [x] Reorganize `README.md` to use a gallery/table for curves.
- [x] Refactor to `SpaceFilling2D` interface.
- [x] Support multiple `GridType`s (Square, Hex, Triangle).
