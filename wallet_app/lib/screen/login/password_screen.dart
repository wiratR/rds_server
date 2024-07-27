import 'package:flutter/material.dart';
import '../../components/appbar/custom_app_bar.dart';
import '../../components/bottomsheet/forgot_password_sheet.dart';
import '../../constants.dart';
// import '../../routes/app_routes.dart';

class PasswordScreen extends StatefulWidget {
  final String phoneNumber; // Add this line to accept the phone number

  PasswordScreen({required this.phoneNumber}); // Update the constructor

  @override
  _PasswordScreenState createState() => _PasswordScreenState();
}

class _PasswordScreenState extends State<PasswordScreen> {
  final TextEditingController _passwordController = TextEditingController();
  bool _isPasswordVisible = false;

  // void _sendOtp() {
  //   // Implement your OTP sending logic here
  //   print('Sending OTP to: ${widget.phoneNumber}');
  //   // Navigate to the OTP verification screen
  //   Navigator.pushNamed(
  //     context,
  //     '/otp-verification',
  //     arguments: OtpVerificationArguments(widget.phoneNumber),
  //   );
  // }

  void _login() {
    // Add your login logic here
    // After login, send OTP
    // _sendOtp();
    print('login with phone : ${widget.phoneNumber}');
    if ((widget.phoneNumber == '+66809921372') &
        (_passwordController.text == '12345678')) {
      print('login conplete');
      Navigator.pushNamed(context, '/home');
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
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: <Widget>[
            Center(
              child: Column(
                children: [
                  Image.asset(
                    'assets/pages/shield_lock_icon.png', // Replace with your icon asset
                    height: 140,
                  ),
                  SizedBox(height: 20),
                  Text(
                    'Enter your password',
                    style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                  ),
                ],
              ),
            ),
            SizedBox(height: 16),
            TextField(
              controller: _passwordController,
              obscureText: !_isPasswordVisible,
              decoration: InputDecoration(
                labelText: 'Password',
                suffixIcon: IconButton(
                  icon: Icon(
                    _isPasswordVisible
                        ? Icons.visibility
                        : Icons.visibility_off,
                  ),
                  onPressed: () {
                    setState(() {
                      _isPasswordVisible = !_isPasswordVisible;
                    });
                  },
                ),
              ),
            ),
            SizedBox(height: 16),
            Align(
              alignment: Alignment.centerRight,
              child: TextButton(
                onPressed: () {
                  showModalBottomSheet(
                    context: context,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.vertical(
                        top: Radius.circular(16.0),
                      ),
                    ),
                    isScrollControlled: true,
                    builder: (context) => ForgotPasswordSheet(),
                  );
                },
                child: Text(
                  'Forgot password?',
                  style: TextStyle(color: Colors.blue),
                ),
              ),
            ),
            SizedBox(height: 32),
            ElevatedButton(
              onPressed: _login,
              child: Text('Login'),
              style: ElevatedButton.styleFrom(
                backgroundColor: kPrimaryColor, // Replace with your color
                minimumSize: Size(double.infinity, 50),
                foregroundColor: Colors.white,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
