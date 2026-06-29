package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

const (
	rsaPK = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

	ecdsaPK = `-----BEGIN PRIVATE KEY-----
MIHuAgEAMBAGByqGSM49AgEGBSuBBAAjBIHWMIHTAgEBBEIAAIWYl3Sv5NZ68ImZ
l0v02Dd2QTThk1hihSPKX37Mv8aHwvnPuAOsDmD0fWRjLWLbPndqDytXq0tdAhcn
ErLTvtahgYkDgYYABAGO9Ki1Xw6LsfrqTv0f11ptkM3lqBcmNKskzQLbkunbFoff
B+qptZDWWccnFsHGoioICOd7NaDXsfN1EOmqwfxXuwEjAKVOw7ewYpsl2EmYGpn4
KdttyyLDZT26zU2/aZG4bBcYXuknPIby7oMX77atn1yOUqPDlq86N4Dsl+XJqJRp
iw==
-----END PRIVATE KEY-----`
)

// TestNewJWT is a unit test for NewJWT
func TestNewJWT(t *testing.T) {
	type args struct {
		cfg JwtConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *JWT
		wantErr bool
	}{
		{
			name: "RSA",
			args: args{
				cfg: JwtConfig{
					Alg:        "RS256",
					PrivateKey: rsaPK,
					Expire:     1,
					RefExpire:  1,
				},
			},
			want: &JWT{
				alg: "RS256",
			},
			wantErr: false,
		},
		{
			name: "ECDSA",
			args: args{
				cfg: JwtConfig{
					Alg:        "ES256",
					PrivateKey: ecdsaPK,
					Expire:     1,
					RefExpire:  1,
				},
			},
			want: &JWT{
				alg: "ES256",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJWT(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.alg != tt.want.alg {
				t.Errorf("NewJWT() got = %v, want %v", got.alg, tt.want.alg)
			}
		})
	}
}

// SignToken is a unit test for SignToken
func TestSignToken(t *testing.T) {
	type args struct {
		cfg    JwtConfig
		claims interface{}
	}
	tests := []struct {
		name    string
		j       *JWT
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "RSA",
			j: &JWT{
				alg:        "RS256",
				privateKey: []byte(rsaPK),
			},
			args: args{
				cfg: JwtConfig{
					Alg:        "RS256",
					PrivateKey: rsaPK,
					Expire:     1,
					RefExpire:  1,
				},
				claims: &map[string]interface{}{
					"sub":  "1234567890",
					"name": "John Doe",
					"iat":  1516239022,
				},
			},
			want:    "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSIsInN1YiI6IjEyMzQ1Njc4OTAifQ.A8Uk-9s_jNdwHd_iaao6Xc6vxw9KhmCDBTSL4olSZM9vWKeIy5cwCB40i_IsOeLk49kKQD1VlEdpxgmNJqnKKflTdWCyiBDqT9HDJGFelh_6sPUBFDByVSkZh4TNKOWTnqbUQjwkpwLmlCtHUp5PV0VQ-rTkqOp4FVsSayO4Qy9qyvMMbldr09IILFfKrBzMUveNyyHncHM_gAUD85lMWqG79ENBzWKjMRLBZbzQ9qLZnXrjbRptcLSE6MkD8OBCciF3_H5_6LV9wuhnQofjzzQfIs-ACONm0hOU-RcSbcOCo_vqUF7_y8zOaWdONuZsxoAMP5m9bWM855e11E-fDK0EUfL8eENnE1EYorNezl63Vhb4XVO-6GNKcPF7X9CFqGN25B0rssKwgbNfdIVTua167NcQ6Bn1ukoGLPMBuHlXzVyegNfbdWwouZ1MWJx5sDmx6BkPuaG1bOsl_jIm5L6Yp8e2m_6za-_jzpeHFKqwZm_bP_N1waVl8fgEJgEER_W98hcdr7EVBMaun_8X8KSOHzHOSU_D1f8nXUzM0qn6AY3BztQe_ijFGz5MHs3va0s9DxKLUMZMaBf0e4dA4c2wC1pXkMiumgThjQL3WNCCxgF8_Fxly1xqc4CPK-50e7qnvmtrhBNXEzQC3c-NTJL65EzI5yeH9oTe8dwBtHY",
			wantErr: false,
		},
		{
			name: "ECDSA",
			j: &JWT{
				alg:        "ES512",
				privateKey: []byte(ecdsaPK),
			},
			args: args{
				cfg: JwtConfig{
					Alg:        "ES512",
					PrivateKey: ecdsaPK,
					Expire:     1,
					RefExpire:  1,
				},
				claims: &map[string]interface{}{
					"sub":  "1234567890",
					"name": "John Doe",
					"iat":  1516239022,
				},
			},
			want:    "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsIm5hbWUiOiJKb2huIERvZSIsInN1YiI6IjEyMzQ1Njc4OTAifQ.MeQ8WOa2VIMx0Wug78vHj9UEbpytSyTJMEqn3gTU1bGiXI3RIQB_Z9K329ezj4K2E0quNbb2dX68yUn35nHwew",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// new jwt
			j, err := NewJWT(tt.args.cfg)
			if err != nil {
				t.Errorf("NewJWT() error = %v", err)
				return
			}

			// sign token
			got, err := j.SignToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// ECDSA signatures are non-deterministic (a random k per signature),
			// so an exact-bytes comparison is invalid. Verify structure and the
			// header algorithm instead. Deterministic algs (RSA PKCS#1 v1.5)
			// still assert exact equality.
			if tt.j.alg == "ES256" || tt.j.alg == "ES384" || tt.j.alg == "ES512" {
				parts := strings.Split(got, ".")
				if len(parts) != 3 {
					t.Fatalf("SignToken() got malformed token: %q", got)
				}
				headerJSON, derr := base64.RawURLEncoding.DecodeString(parts[0])
				if derr != nil {
					t.Fatalf("decode header: %v", derr)
				}
				var header struct {
					Alg string `json:"alg"`
				}
				if uerr := json.Unmarshal(headerJSON, &header); uerr != nil {
					t.Fatalf("unmarshal header: %v", uerr)
				}
				if header.Alg != tt.j.alg {
					t.Errorf("SignToken() alg = %q, want %q", header.Alg, tt.j.alg)
				}
				return
			}

			if got != tt.want {
				t.Errorf("SignToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestParseClaims is a unit test for ParseClaims
func TestParseClaims(t *testing.T) {
	type args struct {
		token  string
		claims interface{}
	}
	tests := []struct {
		name    string
		j       *JWT
		args    args
		wantErr bool
	}{
		{
			name: "RSA",
			j: &JWT{
				alg:        "RS256",
				privateKey: []byte(rsaPK),
			},
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				claims: &map[string]interface{}{
					"sub":  "1234567890",
					"name": "John Doe",
					"iat":  1516239022,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// parse claims
			err := ParseClaims(tt.args.token, tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
