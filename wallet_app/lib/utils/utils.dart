import 'dart:convert'; // For utf8 encoding and hex conversion
import 'package:date_format/date_format.dart';
import 'package:flutter/services.dart' show Uint8List, rootBundle;

/// Custom function to convert bytes to a hexadecimal string
String bytesToHex(List<int> bytes) {
  final buffer = StringBuffer();
  for (var byte in bytes) {
    buffer.write(byte.toRadixString(16).padLeft(2, '0'));
  }
  return buffer.toString();
}

Uint8List hexDecode(String hexString) {
  // Ensure the hex string has an even length
  if (hexString.length % 2 != 0) {
    throw FormatException('Invalid hexadecimal string, length must be even.');
  }

  // Decode hex string to bytes
  List<int> bytes = [];
  for (int i = 0; i < hexString.length; i += 2) {
    String byteString = hexString.substring(i, i + 2);
    int byte = int.parse(byteString, radix: 16);
    bytes.add(byte);
  }

  return Uint8List.fromList(bytes);
}

String concatJson(Map<String, String> jsonData) {
  // Sort and concatenate the parameters
  var keys = jsonData.keys.where((key) => key != 'sign').toList();
  keys.sort();
  var strJson = keys.map((key) => '$key=${jsonData[key]}').join('');
  return strJson;
}

String generateTimeStamp() {
  final DateTime now = DateTime.now();
  return formatDate(now, [yyyy, mm, dd, HH, nn, ss]);
}

String generateMchOrderNo() {
  final DateTime now = DateTime.now();
  final DateTime oneDayAgo = now.subtract(Duration(days: 1));
  return formatDate(oneDayAgo, [yyyy, mm, dd, HH, nn, ss]);
}

Future<String> loadPrivateKey() async {
  return await rootBundle.loadString('assets/key/Mch38806_PrivateKey.pem');
}

String uint8ListToString(Uint8List uint8List) {
  return utf8.decode(uint8List);
}

String parsePEM(String pem) {
  return pem
      .replaceAll('-----BEGIN RSA PRIVATE KEY-----', '')
      .replaceAll('-----END RSA PRIVATE KEY-----', '')
      .replaceAll('\n', '');
}

String parsePEMPublic(String pem) {
  return pem
      .replaceAll('-----BEGIN PUBLIC KEY-----', '')
      .replaceAll('-----END PUBLIC KEY-----', '')
      .replaceAll('\n', '');
}
