GO = go
ANDROID_SDK_ROOT = ""
ANDROID_JAVA_ROOT = ""
TEMP = ""

ifeq ($(shell uname -s), Darwin)
    ANDROID_SDK_ROOT = "$(HOME)/Library/Android/sdk/"
    ANDROID_JAVA_ROOT = /Library/Java/JavaVirtualMachines/zulu-11.jdk/Contents/Home/bin
    ANDROID_PLATFORM = $(ANDROID_SDK_ROOT)/platforms/$(shell ls $(ANDROID_SDK_ROOT)/platforms | sort -n | tail -n 1)
    TEMP = /tmp
endif

macos:
	$(GO) run gioui.org/cmd/gogio -schemes "testings" -target macos -arch arm64 -o app.app .

windows:
	$(GO) run gioui.org/cmd/gogio -schemes "testings" -target windows -arch amd64 -o app.exe .

ios:
	$(GO) run gioui.org/cmd/gogio -schemes "testings" -target ios -arch amd64 -o ios.app .

android:
	ANDROID_SDK_ROOT=$(ANDROID_SDK_ROOT) PATH=$(ANDROID_JAVA_ROOT):$(PATH) $(GO) run gioui.org/cmd/gogio -schemes "testings" -target android -arch arm64 -o app.apk .