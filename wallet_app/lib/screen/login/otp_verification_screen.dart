import 'package:flutter/material.dart';

import '../../components/appbar/custom_app_bar.dart';
import '../../constants.dart';

class OtpVerificationScreen extends StatefulWidget {
  final String phoneNumber;

  OtpVerificationScreen({required this.phoneNumber});

  @override
  _OtpVerificationScreenState createState() => _OtpVerificationScreenState();
}

class _OtpVerificationScreenState extends State<OtpVerificationScreen> {
  final _formKey = GlobalKey<FormState>();
  String? _otpCode;

  void _verifyOtp() {
    if (_formKey.currentState!.validate()) {
      // Perform OTP verification logic here
      print('Verifying OTP: $_otpCode for phone number: ${widget.phoneNumber}');
      // Navigate to the next screen upon successful OTP verification
      Navigator.pushNamed(context, '/password');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png', // Path to your logo asset
        showBackButton:
            true, // Set this to true if you want to show the back button
        showIconLogout: false,
      ), // Use the custom app bar
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            children: <Widget>[
              SizedBox(height: 60),
              Text(
                'Verify OTP',
                style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 80),
              Form(
                key: _formKey,
                child: Column(
                  children: <Widget>[
                    TextFormField(
                      decoration: InputDecoration(
                        labelText: 'OTP Code',
                        border: OutlineInputBorder(),
                      ),
                      keyboardType: TextInputType.number,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter the OTP code';
                        }
                        return null;
                      },
                      onChanged: (value) {
                        _otpCode = value;
                      },
                    ),
                    SizedBox(height: 20),
                    ElevatedButton(
                      onPressed: _verifyOtp,
                      style: ElevatedButton.styleFrom(
                        backgroundColor: kPrimaryColor,
                        minimumSize: Size(double.infinity, 50),
                        foregroundColor: Colors.white,
                      ),
                      child: Text('Verify'),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
