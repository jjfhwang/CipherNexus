# CipherNexus: Secure and Flexible Encryption Library

CipherNexus is a Go library designed to provide developers with a robust and versatile set of cryptographic tools. It aims to simplify the implementation of secure data handling practices within Go applications by offering a clean, well-documented, and performant interface to various encryption algorithms and utilities. This library is intended for developers who require a secure and reliable means to encrypt and decrypt sensitive data, manage cryptographic keys, and implement advanced security protocols within their Go projects.

CipherNexus addresses the growing need for accessible and easily integrable encryption solutions. Many existing cryptographic libraries can be complex and challenging to implement correctly, increasing the risk of security vulnerabilities due to misconfiguration or improper usage. CipherNexus aims to mitigate these risks by providing a higher-level abstraction that simplifies common encryption tasks while still allowing for fine-grained control over cryptographic parameters. The library prioritizes security best practices and incorporates features to help developers avoid common pitfalls in cryptographic implementation.

The primary focus of CipherNexus is to provide a flexible and customizable encryption framework. It supports a range of symmetric encryption algorithms, including AES and ChaCha20, with various key sizes and operating modes. The library also includes utilities for key derivation, secure random number generation, and message authentication. Furthermore, CipherNexus is designed to be extensible, allowing developers to easily integrate new encryption algorithms or custom security protocols as needed. By providing a comprehensive suite of cryptographic tools, CipherNexus empowers developers to build secure and reliable applications with confidence.

CipherNexus offers several benefits compared to directly using lower-level cryptographic primitives. It provides a consistent and intuitive API, which simplifies the process of encryption and decryption. It handles many of the complexities of cryptographic implementation, such as padding, initialization vector management, and error handling, reducing the likelihood of introducing vulnerabilities. Moreover, CipherNexus is designed to be performant, leveraging Go's built-in cryptographic primitives and optimized for speed and efficiency. The library also includes comprehensive documentation and examples to help developers get started quickly and use the library correctly.

## Key Features

*   **AES Encryption:** Supports Advanced Encryption Standard (AES) with key sizes of 128, 192, and 256 bits. Implements Cipher Feedback (CFB), Output Feedback (OFB), Counter (CTR), and Galois/Counter Mode (GCM) modes of operation for flexibility in different security scenarios. The GCM implementation leverages authenticated encryption to provide both confidentiality and integrity.

*   **ChaCha20 Encryption:** Provides support for the ChaCha20 stream cipher, a modern and high-performance alternative to AES. Includes the Poly1305 authenticator when using ChaCha20, offering authenticated encryption for data integrity verification.

*   **Key Derivation:** Implements Password-Based Key Derivation Function 2 (PBKDF2) with SHA-256 and SHA-512 hashing algorithms. Allows for secure key derivation from passwords using salting and iteration to enhance security against brute-force attacks. Supports customizable salt lengths and iteration counts.

*   **Secure Random Number Generation:** Uses Go's `crypto/rand` package to generate cryptographically secure random numbers. Provides functions for generating random bytes and random strings for use as salts, initialization vectors, and cryptographic keys.

*   **Message Authentication Code (MAC):** Implements HMAC (Hash-based Message Authentication Code) using SHA-256 for data integrity verification. Allows for the creation of authentication tags that can be used to ensure that data has not been tampered with during transmission or storage.

*   **Error Handling:** Provides comprehensive error handling throughout the library. Returns specific error types to allow developers to handle errors gracefully and prevent security vulnerabilities due to unexpected errors.

*   **Extensible Architecture:** Designed to be extensible, allowing developers to add support for new encryption algorithms, key derivation functions, and authentication methods. The library provides interfaces and abstract classes that make it easy to integrate custom cryptographic components.

## Technology Stack

*   **Go:** The primary programming language used for the library. Go's performance and concurrency features make it well-suited for cryptographic applications.
*   **crypto/aes:** Go's built-in package for AES encryption. CipherNexus leverages this package for its AES implementation.
*   **crypto/cipher:** Go's built-in package for block cipher modes of operation. CipherNexus utilizes this package for implementing CFB, OFB, CTR, and GCM modes.
*   **crypto/rand:** Go's built-in package for generating cryptographically secure random numbers. CipherNexus uses this package for generating salts, initialization vectors, and keys.
*   **crypto/sha256 & crypto/sha512:** Go's built-in packages for SHA-256 and SHA-512 hashing algorithms. CipherNexus uses these packages for key derivation and HMAC.
*   **golang.org/x/crypto/chacha20poly1305:** External Go package for ChaCha20-Poly1305 authenticated encryption. CipherNexus utilizes this package for ChaCha20 encryption with authentication.

## Installation

To install CipherNexus, use the following `go get` command:

go get github.com/jjfhwang/CipherNexus

This command will download and install the library and its dependencies into your Go workspace.

## Configuration

CipherNexus can be configured through environment variables or directly within your Go code.

*   **Key Length:** The length of the encryption key can be specified directly in the function calls. For example, when using AES, you can choose between 128, 192, or 256-bit keys.
*   **Iteration Count:** The iteration count for PBKDF2 can be adjusted to increase the computational cost of key derivation and improve security against brute-force attacks.
*   **Salt Length:** The length of the salt used in PBKDF2 can be specified. A longer salt improves security by making rainbow table attacks more difficult.

## Usage

Here are some examples of how to use CipherNexus:

Encrypting and decrypting data with AES-256-GCM:



Deriving a key from a password using PBKDF2:



Detailed API documentation will be available on the project's GitHub Pages site soon.

## Contributing

We welcome contributions to CipherNexus! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with appropriate comments.
4.  Include unit tests to verify the correctness of your changes.
5.  Submit a pull request with a detailed description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/CipherNexus/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to acknowledge the developers of the Go programming language and the authors of the various cryptographic libraries that CipherNexus relies upon. Their work has made this project possible.