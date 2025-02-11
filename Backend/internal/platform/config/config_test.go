package config

import (
	"reflect"
	"testing"
)

func TestGetDBUserName(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "Success - GetDBUserName Set Variable",
			want: "Test",
		},
		{
			name: "Success - GetDBUserName Unset Variable",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DB_USERNAME", tt.want)
			got := GetDBUserName()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBUserName got = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetDBPassword(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "Success - GetDBPassword Set Variable",
			want: "Test",
		},
		{
			name: "Success - GetDBPassword Unset Variable",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DB_PASSWORD", tt.want)
			got := GetDBPassword()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBPassword got = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetDBHost(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "Success - GetDBHost Set Variable",
			want: "Test",
		},
		{
			name: "Success - GetDBHost Unset Variable",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DB_HOST", tt.want)
			got := GetDBHost()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBHost got = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetDBName(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "Success - GetDBName Set Variable",
			want: "Test",
		},
		{
			name: "Success - GetDBName Unset Variable",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DB_NAME", tt.want)
			got := GetDBName()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBName got = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetServerHost(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "Success - GetServerHost Set Variable",
			want: "Test",
		},
		{
			name: "Success - GetServerHost Unset Variable",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BACKEND_HOST", tt.want)
			got := GetServerHost()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServerHost got = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetAllowedHosts(t *testing.T) {
	tests := []struct {
		name string

		want string
	}{
		{
			name: "Success - GetAllowedHosts Set Variable",
			want: "Test",
		},
		{
			name: "Success - GetAllowedHosts Unset Variable",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BACKEND_ALLOWED_HOSTS", tt.want)
			got := GetAllowedHosts()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllowedHosts got = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestGetProductionFlag(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want bool
	}{
		{
			name: "Success -  Set True",
			set:  "true",
			want: true,
		},
		{
			name: "Success -  Unset",
			set:  "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BACKEND_PROD_FLAG", tt.set)
			got := GetProductionFlag()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductionFlag got = %v, want: %v", got, tt.want)
			}
		})
	}
}
