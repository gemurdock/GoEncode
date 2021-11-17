package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testText struct {
	text      string
	shouldErr bool
}

func TestEncodeDecode(t *testing.T) {
	assert := assert.New(t)

	testTexts := []testText{
		{"Hello World!", false},
		{"How are you doing today?", false},
		{"k2YBXm}e*;,xq>A\"8&kT", false},
		{"Zn/'n7uH%$8(QC2+=&AE", false},
		{"", true},
		{"SaP)[d(]3DpvTb9qpCC=B?*g<H;WDzt+", false}, // 32 chars
		{"'\\,[RH3.?49VA?G#2<CM5K2CLsdE:pXYeuPMS[uxT)(*7fEg:XKM[<6Bb9X;-^z;", false}, // 64 chars,
		{`5P<..&4#dfUEUdbjt$,9!}XW$y-wmCY79KDFg?[b_LyFFzGP<keU~3@#(4/<>FJM&#)$SZ7D]=Z
		\rb?5=@!x-kmJ;%r5SV=>kPW-?dLh?p9dqWY\\DgphtzRb_7WEkHhm-`, false}, // 128 chars
		{`zRAKbhB6FFd*3vNzB$9hn^q9K&Mx+B=fjg9=DX#x2rLZNqVD-#P_yyGf2cxS**Rd_&m-qnh6?*3
		%8af3ns$Y6twk4efjV?YUPYwsjp7Rd&V-F6jndYm9_9e%Q@N%rGPHGStZX7kgHJsSBtt-*yPp@St#
		fpd_q6yDbnLeCd7h&zN#8UW382D7TbKd-yLZyJg37CqY%HUkGFeZ$mBja_hPz5SamA+-sb=5PcC&c
		=%YQ74$GKg*6px5^UcnHF@^vjTj`, false}, // 256 chars
		{`sh$SK?bDxSVE-CE!N28Xg8P+uGpm$$%jb4*qTpy36j$v=dX=gA64#fp2Ev5xtMVcaw9hRc6*$4N
		x&MA&-M4Xwq4?2b=87=d$A8vq_WxB8#RY-8jc!uqZ3&nrK&%!cthUkWx3KRq9CUm!MfQ9+AhGtAVz
		qWX8JC!+8H3t?BN+Z5=Yf74pCrR^WC3cMSu#bVPguUuL-rue-$b6Nvrk&WFsP5q%N5UBnk2dej5pa
		_^UJ$sh#YQT9*BUS_7JUv_7T?NmJYjPM$A4#f#aDRFJJ&aJYA7cA3&Bj#^V_q6XE7t9rFFLYD7gze
		J=@BbdSd?YtKCJt7M#F9usWyKe3qC7KAg__2=tXdJ9=h7QUKmWyz*X?Z-v@#fD=4HhYxa!AUQ@ufY
		WAqWvch&hNWkzKg3NGg!H4cBDg+EgaDv@a!$8+q#8HXHPwJ$vcDjrv^Pvc2EDD3UR#2bqvUa&R$ZM
		%N@h4dwwz7p$5V@3&NZehFna2_!Z8T4ZFF#4bfcx6*H9sr4RxRk=`, false}, // 512 chars
	}
	testPasswords := []testText{
		{"password", false},
		{"a", false},
		{"ab", false},
		{"abc", false},
		{"xfkoybvd", false},
		{"rK^\\8)rY8VQ$yP&R~4SZ", false},
		{"", true},
		{"4Sx?BU2p", false},
		{"CNBnKt&Be@4?j4dj", false},
		{"a*BAq8WDELS&@Y#rmBt^-g$v+634qTEZ", false},
		{"v2tX#2EVby5r6!%5LftDqNUTjw?J7?*jq6rg$4v62Abj&Bpd_2F-?7%jkmeRCqSG", false},
	}

	for _, testText := range testTexts {
		for _, testPassword := range testPasswords {
			encoded, err := Encode(testText.text, testPassword.text)
			if testText.shouldErr || testPassword.shouldErr {
				assert.NotEqual(err, nil)
			} else {
				assert.Equal(err, nil)
				assert.NotEqual(testText.text, encoded)

				decoded, err := Decode(encoded, testPassword.text)
				assert.Equal(err, nil)
				assert.Equal(testText.text, decoded)
			}
		}
	}
}
