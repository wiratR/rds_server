import 'package:flutter/material.dart';
import '../../constants.dart';
import '../../routes/app_routes.dart';

// Temporary replacement for the removed PhoneNumber class from intl_phone_number_input
class PhoneNumber {
  final String? phoneNumber;
  final String isoCode;

  PhoneNumber({required this.isoCode, this.phoneNumber});
}

class LoginScreen extends StatefulWidget {
  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _formKey = GlobalKey<FormState>();
  final _phoneController = TextEditingController();

  PhoneNumber number = PhoneNumber(isoCode: 'TH');

  void _submit() {
    if (_formKey.currentState!.validate()) {
      number = PhoneNumber(
        isoCode: 'TH',
        phoneNumber: _phoneController.text,
      );

      debugPrint('Phone number: ${number.phoneNumber}');
      Navigator.pushNamed(
        context,
        '/password',
        arguments: PasswordScreenArguments(number.phoneNumber!),
      );
    }
  }

  @override
  void dispose() {
    _phoneController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            children: <Widget>[
              const SizedBox(height: 60),
              Image.asset('assets/logo.png', height: 60),
              const SizedBox(height: 20),
              Image.asset('assets/pages/phone.png', height: 220),
              const SizedBox(height: 20),
              const Text(
                'Enter your mobile number',
                style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 50),
              Form(
                key: _formKey,
                child: Column(
                  children: <Widget>[
                    TextFormField(
                      controller: _phoneController,
                      keyboardType: TextInputType.phone,
                      decoration: const InputDecoration(
                        labelText: 'Mobile number',
                        border: OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter a valid phone number';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 20),
                    ElevatedButton(
                      onPressed: _submit,
                      style: ElevatedButton.styleFrom(
                        backgroundColor: kPrimaryColor,
                        minimumSize: const Size(double.infinity, 50),
                        foregroundColor: Colors.white,
                      ),
                      child: const Text('Continue'),
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 20),
              const Text('or continue using'),
              const SizedBox(height: 20),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: <Widget>[
                  IconButton(
                    icon: Image.asset('assets/icons/facebook.png'),
                    onPressed: () {},
                  ),
                  IconButton(
                    icon: Image.asset('assets/icons/google.png'),
                    onPressed: () {},
                  ),
                  IconButton(
                    icon: Image.asset('assets/icons/apple.png'),
                    onPressed: () {},
                  ),
                ],
              ),
              const SizedBox(height: 10),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text('Don\'t have an account?'),
                  TextButton(
                    onPressed: () {
                      Navigator.pushNamed(context, '/register');
                    },
                    child: const Text('Create New'),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
