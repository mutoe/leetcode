package validate_ip_address

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "172.16.254.1",
			args: args{"172.16.254.1"},
			want: "IPv4",
		},
		{
			name: "01.01.01.01",
			args: args{"01.01.01.01"},
			want: "Neither",
		},
		{
			name: "192.168.1.0",
			args: args{"192.168.1.0"},
			want: "IPv4",
		},
		{
			name: "1.0.1.",
			args: args{"1.0.1."},
			want: "Neither",
		},
		{
			name: "2001:0db8:85a3:0:0:8A2E:0370:7334",
			args: args{"2001:0db8:85a3:0:0:8A2E:0370:7334"},
			want: "IPv6",
		},
		{
			name: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			args: args{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			want: "IPv6",
		},
		{
			name: "256.256.256.256",
			args: args{"256.256.256.256"},
			want: "Neither",
		},
		{
			name: "2001:db8:85a3:0::8a2E:0370:7334",
			args: args{"2001:db8:85a3:0::8a2E:0370:7334"},
			want: "Neither",
		},
		{
			name: "2001:0db8:85a3:0:0:8A2E:0370:7334:",
			args: args{"2001:0db8:85a3:0:0:8A2E:0370:7334:"},
			want: "Neither",
		},
		{
			name: "1e1.4.5.6",
			args: args{"1e1.4.5.6"},
			want: "Neither",
		},
		{
			name: "2001:0db8:85a3::8A2E:037j:7334:",
			args: args{"2001:0db8:85a3::8A2E:037j:7334:"},
			want: "Neither",
		},
		{
			name: "2001:0db8:85a3::8A2E:037j:7334:",
			args: args{"2001:0db8:85a3:0:8A2E:037j:7334:0"},
			want: "Neither",
		},
		{
			name: "02001:0db8:85a3:0000:0000:8a2e:0370:7334",
			args: args{"02001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			want: "Neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
