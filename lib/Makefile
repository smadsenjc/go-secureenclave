libSEP256.a: libSEP256-arm64.a libSEP256-x86_64.a
	lipo -create $^ -o $@

libSEP256-arm64.a: SEP256.swift
	xcrun -sdk macosx swiftc $^ -module-name SEP256 -target arm64-apple-macos10.15 -framework CryptoKit -emit-library -static -o $@

libSEP256-x86_64.a: SEP256.swift
	xcrun -sdk macosx swiftc $^ -module-name SEP256 -target x86_64-apple-macos10.15 -framework CryptoKit -emit-library -static -o $@

clean:
	rm -f libSEP256.a libSEP256-arm64.a libSEP256-x86_64.a
