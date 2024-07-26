import 'dart:convert';
import 'dart:typed_data';
import 'package:flutter/material.dart';
import '../../../../../components/appbar/custom_app_bar.dart';
import '../../../../../service/ksher/ksherpay_service.dart';

class PromptpayScreen extends StatelessWidget {
  final KsherResp responseMessage;
  PromptpayScreen({required this.responseMessage});

  @override
  Widget build(BuildContext context) {
    Map<String, String> stringMap = responseMessage.data
        .map((key, value) => MapEntry(key, value.toString()));
    String base64String = stringMap['imgdat']!.split(',').last;
    // Decode base64 string to Uint8List
    Uint8List bytes = base64Decode(base64String);

    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png', // Path to your logo asset
        showBackButton: true,
        showIconLogout: false,
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            // Display the QR code image
            Image.memory(bytes),
            SizedBox(height: 20),
            Text(
              'This is Promtpay screen.',
              style: TextStyle(fontSize: 24),
            ),
          ],
        ),
      ),
    );
  }
}
