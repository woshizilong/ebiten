// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opengl

type Image struct {
	Texture     Texture
	Framebuffer *Framebuffer
	width       int
	height      int
}

func NewImage(width, height int) *Image {
	return &Image{
		width:  width,
		height: height,
	}
}

func (i *Image) Size() (int, int) {
	return i.width, i.height
}

func (i *Image) IsInvalidated() bool {
	return !theContext.isTexture(i.Texture)
}

func (i *Image) Delete() {
	if i.Framebuffer != nil {
		i.Framebuffer.Delete()
	}
	if i.Texture != *new(Texture) {
		theContext.deleteTexture(i.Texture)
	}
}