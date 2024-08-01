import 'package:flutter/material.dart';
import '../../components/appbar/custom_app_bar.dart';
import '../../components/bottomsheet/forgot_password_sheet.dart';
import '../../constants.dart';
import '../../service/auth/auth_service.dart';
import '../main/home/home_screen.dart';

class PasswordScreen extends StatefulWidget {
  final String phoneNumber; // Add this line to accept the phone number

  const PasswordScreen(
      {super.key, required this.phoneNumber}); // Update the constructor

  @override
  // ignore: library_private_types_in_public_api
  _PasswordScreenState createState() => _PasswordScreenState();
}

class _PasswordScreenState extends State<PasswordScreen> {
  final AuthService authService = AuthService();
  final TextEditingController _passwordController = TextEditingController();
  bool _isPasswordVisible = false;

  // void _sendOtp() {
  //   // Implement your OTP sending logic here
  //   debugPrint('Sending OTP to: ${widget.phoneNumber}');
  //   // Navigate to the OTP verification screen
  //   Navigator.pushNamed(
  //     context,
  //     '/otp-verification',
  //     arguments: OtpVerificationArguments(widget.phoneNumber),
  //   );
  // }

  Future<void> _login() async {
    // Add your login logic here
    // After login, send OTP
    // _sendOtp();
    debugPrint('login with phone : ${widget.phoneNumber}');
    debugPrint('login with password : ${_passwordController.text}');

    String? token = await authService.loginByPhone(
      widget.phoneNumber,
      _passwordController.text,
    );

    if (token != null) {
      // Navigate to next screen or handle successful login
      debugPrint('Login successful! Token: $token');
      // Navigate to the home screen on success
      Navigator.pushReplacement(
        // ignore: use_build_context_synchronously
        context,
        MaterialPageRoute(builder: (context) => HomeScreen()),
      );
    } else {
      // Show error message or handle login failure
      debugPrint('Login failed');
      // ignore: use_build_context_synchronously
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text('Login failed. Please try again.'),
          duration: Duration(seconds: 3),
        ),
      );
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
                  const SizedBox(height: 20),
                  const Text(
                    'Enter your password',
                    style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 16),
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
            const SizedBox(height: 16),
            Align(
              alignment: Alignment.centerRight,
              child: TextButton(
                onPressed: () {
                  showModalBottomSheet(
                    context: context,
                    shape: const RoundedRectangleBorder(
                      borderRadius: BorderRadius.vertical(
                        top: Radius.circular(16.0),
                      ),
                    ),
                    isScrollControlled: true,
                    builder: (context) => const ForgotPasswordSheet(),
                  );
                },
                child: const Text(
                  'Forgot password?',
                  style: TextStyle(color: Colors.blue),
                ),
              ),
            ),
            const SizedBox(height: 32),
            ElevatedButton(
              onPressed: _login,
              style: ElevatedButton.styleFrom(
                backgroundColor: kPrimaryColor, // Replace with your color
                minimumSize: const Size(double.infinity, 50),
                foregroundColor: Colors.white,
              ),
              child: const Text('Login'),
            ),
          ],
        ),
      ),
    );
  }
}
