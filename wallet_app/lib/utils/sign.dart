import 'dart:convert';
import 'package:fast_rsa/fast_rsa.dart';
import 'package:flutter/foundation.dart';
import 'utils.dart';

Future<void> signAndVerify(
    String stgTosign, String privateKeyPem, String publicKeyPem) async {
  var signature = await RSA.signPKCS1v15(stgTosign, Hash.MD5, privateKeyPem);
  // Convert the signature to bytes
  List<int> signatureBytes = base64.decode(signature);
  // Convert the bytes to a hexadecimal string using the custom function
  String hexSignature = bytesToHex(signatureBytes);
  debugPrint("Data after sign (Hex) = $hexSignature\n");

  // var isVerified = await RSA.verifyPSS(
  //     signature, stgTosign, Hash.MD5, SaltLength.AUTO, publicKeyPem);

  // debugPrint("verify result = $isVerified\n");
}
