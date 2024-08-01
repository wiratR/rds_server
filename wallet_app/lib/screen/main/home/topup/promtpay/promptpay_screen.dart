import 'dart:convert';
import 'dart:typed_data';
import 'package:flutter/material.dart';
import '../../../../../components/appbar/custom_app_bar.dart';
import '../../../../../service/ksher/ksherpay_service.dart';

class PromptpayScreen extends StatelessWidget {
  final KsherResp responseMessage;
  const PromptpayScreen({super.key, required this.responseMessage});

  @override
  Widget build(BuildContext context) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _showResponseDialog(context);
    });

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
            const SizedBox(height: 20),
            const Text(
              'This is Promptpay screen.',
              style: TextStyle(fontSize: 24),
            ),
          ],
        ),
      ),
    );
  }

  void _showResponseDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Payment Response'),
          content: Text('Response: ${responseMessage.toJson()}'),
          actions: [
            TextButton(
              child: const Text('OK'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }
}
