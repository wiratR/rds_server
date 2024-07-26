import 'dart:async';
import 'dart:math';
import 'package:flutter/material.dart';
import 'package:qr_flutter/qr_flutter.dart';

class CardScreen extends StatefulWidget {
  @override
  _CardScreenState createState() => _CardScreenState();
}

class _CardScreenState extends State<CardScreen> {
  late String _qrData;
  late Timer _timer;
  int _countdown = 60;

  @override
  void initState() {
    super.initState();
    _generateQRCode();
    _startTimer();
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  void _generateQRCode() {
    final random = Random();
    _qrData = random
        .nextInt(100000)
        .toString(); // Replace with your own logic to generate QR data
  }

  void _startTimer() {
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      setState(() {
        if (_countdown > 0) {
          _countdown--;
        } else {
          _countdown = 60;
          _generateQRCode();
        }
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    // Get the size of the screen
    double screenWidth = MediaQuery.of(context).size.width;
    double screenHeight = MediaQuery.of(context).size.height;
    // Determine the size of the QR code based on the screen size
    double qrSize = min(screenWidth, screenHeight) * 0.7;

    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            if (_qrData.isNotEmpty)
              QrImageView(
                data: _qrData,
                version: QrVersions.auto,
                size: qrSize,
              ),
            SizedBox(height: 20),
            Text(
              'Plase Tap this QR pn GATE',
              style: TextStyle(fontSize: 16),
            ),
            SizedBox(height: 20),
            Text(
              'Refreshing in $_countdown seconds',
              style: TextStyle(fontSize: 16),
            ),
          ],
        ),
      ),
    );
  }
}
