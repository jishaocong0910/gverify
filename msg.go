/*
Copyright 2024-present jishaocong0910

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vfy

type msgBuildFunc func(f *FieldInfo)

func msgBuildFuncDefault(f *FieldInfo) {
	f.Msg("%s is illegal", f.fieldName)
}

func msgBuildFuncRequired(f *FieldInfo) {
	f.Msg("%s is required", f.fieldName)
}

func msgBuildFuncMin(f *FieldInfo) {
	f.Msg("%s must not be less than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncMax(f *FieldInfo) {
	f.Msg("%s must not be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncRange(f *FieldInfo) {
	f.Msg("%s must be %s to %s", f.fieldName, f.Confine(0), f.Confine(1))
}

func msgBuildFuncGt(f *FieldInfo) {
	f.Msg("%s must be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncLt(f *FieldInfo) {
	f.Msg("%s must be less than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncWithin(f *FieldInfo) {
	f.Msg("%s must be greater than %s and less than %s", f.fieldName, f.Confine(0), f.Confine(1))
}

func msgBuildFuncEnum(f *FieldInfo) {
	f.Msg("%s must be %s", f.fieldName, f.Confines())
}

func msgBuildFuncNotBlank(f *FieldInfo) {
	f.Msg("%s must not be blank", f.fieldName)
}

func msgBuildFuncRegex(f *FieldInfo) {
	f.Msg("%s's format is illegal", f.fieldName)
}

func msgBuildFuncNotEmpty(f *FieldInfo) {
	f.Msg("%s must not be empty", f.fieldName)
}

func msgBuildFuncLength(f *FieldInfo) {
	f.Msg("%s's length must be %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncLengthMin(f *FieldInfo) {
	f.Msg("%s's length must not be less than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncLengthMax(f *FieldInfo) {
	f.Msg("%s's length must not be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncLengthRange(f *FieldInfo) {
	f.Msg("%s's length must be %s to %s", f.fieldName, f.Confine(0), f.Confine(1))
}

func msgBuildFuncLengthGt(f *FieldInfo) {
	f.Msg("%s's length must be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncLengthLt(f *FieldInfo) {
	f.Msg("%s's length must be less than %s", f.fieldName, f.Confine(0))
}

func msgBuildFuncLengthWithin(f *FieldInfo) {
	f.Msg("%s's length must be greater than %s and less than %s", f.fieldName, f.Confine(0), f.Confine(1))
}
