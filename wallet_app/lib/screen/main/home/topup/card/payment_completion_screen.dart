import 'package:flutter/material.dart';
import 'dart:async';
import '../../../../../components/appbar/custom_app_bar.dart';

class PaymentCompletionScreen extends StatefulWidget {
  final String paymentMessage;

  const PaymentCompletionScreen({super.key, required this.paymentMessage});

  @override
  _PaymentCompletionScreenState createState() =>
      _PaymentCompletionScreenState();
}

class _PaymentCompletionScreenState extends State<PaymentCompletionScreen> {
  int _countdown = 10;
  Timer? _timer;

  @override
  void initState() {
    super.initState();
    _startCountdown();
  }

  void _startCountdown() {
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      if (_countdown == 0) {
        _timer?.cancel();
        Navigator.pushNamedAndRemoveUntil(context, '/home', (route) => false);
      } else {
        setState(() {
          _countdown--;
        });
      }
    });
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png',
        showBackButton: true,
        showIconLogout: false,
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Icon(
              Icons.check_circle,
              color: Colors.green,
              size: 100.0,
            ),
            const SizedBox(height: 20.0),
            const Text(
              'Payment Successful!',
              style: TextStyle(fontSize: 24.0, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 20.0),
            Text(
              widget.paymentMessage,
              style: const TextStyle(fontSize: 16.0),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 20.0),
            Text(
              'Returning to Home in $_countdown seconds',
              style: const TextStyle(fontSize: 16.0),
              textAlign: TextAlign.center,
            ),
          ],
        ),
      ),
    );
  }
}
