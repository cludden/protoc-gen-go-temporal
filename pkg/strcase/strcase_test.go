/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strcase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitWords(t *testing.T) {
	c := NewCaser(WithAcronyms("AWS", "API", "PostgresQL", "AbcdAcme"))
	cases := []struct {
		input    string
		expected []string
	}{
		{"test_case", []string{"test", "case"}},
		{"test.case", []string{"test", "case"}},
		{"test", []string{"test"}},
		{"TestCase", []string{"Test", "Case"}},
		{" test  case ", []string{"test", "case"}},
		{"", []string{}},
		{"many_many_words", []string{"many", "many", "words"}},
		{"AnyKind of_string", []string{"Any", "Kind", "of", "string"}},
		{"AnyKind.of-string", []string{"Any", "Kind", "of", "string"}},
		{"odd-fix", []string{"odd", "fix"}},
		{"numbers2And55with000", []string{"numbers", "2", "And", "55", "with", "000"}},
		{"ID", []string{"ID"}},
		{"CONSTANT_CASE", []string{"CONSTANT", "CASE"}},
		{"AWS", []string{"AWS"}},
		{"AWSInstance", []string{"AWS", "Instance"}},
		{"AWSInstanceID", []string{"AWS", "Instance", "ID"}},
		{"ManageAWS", []string{"Manage", "AWS"}},
		{"CreatePostgresQLCluster", []string{"Create", "PostgresQL", "Cluster"}},
	}

	for _, tc := range cases {
		t.Logf("testing %s", tc.input)
		result := c.splitWords(tc.input)
		require.Equal(t, tc.expected, result)
	}
}

func TestToCamel(t *testing.T) {
	c := NewCaser(WithAcronyms("AWS", "API", "PostgresQL", "AbcdAcme"))
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
		{"ID", "Id"},
		{"CONSTANT_CASE", "ConstantCase"},
		{"AWS", "AWS"},
		{"AWSInstance", "AWSInstance"},
		{"AWSInstanceID", "AWSInstanceId"},
		{"ManageAWS", "ManageAWS"},
		{"CreatePostgresQLCluster", "CreatePostgresQLCluster"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := c.ToCamel(in)
		require.Equal(t, out, result)
	}
}

func TestToKebab(t *testing.T) {
	c := NewCaser(WithAcronyms("AWS", "API", "PostgresQL", "AbcdAcme"))
	cases := [][]string{
		{"test_case", "test-case"},
		{"test.case", "test-case"},
		{"test", "test"},
		{"TestCase", "test-case"},
		{" test  case ", "test-case"},
		{"", ""},
		{"many_many_words", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"odd-fix", "odd-fix"},
		{"numbers2And55with000", "numbers-2-and-55-with-000"},
		{"ID", "id"},
		{"CONSTANT_CASE", "constant-case"},
		{"AWS", "aws"},
		{"AWSInstance", "aws-instance"},
		{"AWSInstanceID", "aws-instance-id"},
		{"ManageAWS", "manage-aws"},
		{"CreatePostgresQLCluster", "create-postgresql-cluster"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := c.ToKebab(in)
		require.Equal(t, out, result)
	}
}

func TestToLowerCamel(t *testing.T) {
	c := NewCaser(WithAcronyms("AWS", "API", "PostgresQL", "AbcdAcme"))
	cases := [][]string{
		{"test_case", "testCase"},
		{"test.case", "testCase"},
		{"test", "test"},
		{"TestCase", "testCase"},
		{" test  case ", "testCase"},
		{"", ""},
		{"many_many_words", "manyManyWords"},
		{"AnyKind of_string", "anyKindOfString"},
		{"odd-fix", "oddFix"},
		{"numbers2And55with000", "numbers2And55With000"},
		{"ID", "id"},
		{"CONSTANT_CASE", "constantCase"},
		{"AWS", "aws"},
		{"AWSInstance", "awsInstance"},
		{"AWSInstanceID", "awsInstanceId"},
		{"ManageAWS", "manageAWS"},
		{"CreatePostgresQLCluster", "createPostgresQLCluster"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := c.ToLowerCamel(in)
		require.Equal(t, out, result)
	}
}
