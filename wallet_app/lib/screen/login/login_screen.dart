import 'package:flutter/material.dart';
import 'package:intl_phone_number_input/intl_phone_number_input.dart';
import '../../constants.dart';
import '../../routes/app_routes.dart';

class LoginScreen extends StatefulWidget {
  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _formKey = GlobalKey<FormState>();
  String initialCountry = 'TH';
  PhoneNumber number = PhoneNumber(isoCode: 'TH');

  void _submit() {
    if (_formKey.currentState!.validate()) {
      // Perform login logic
      debugPrint('Phone number: ${number.phoneNumber}');
      // Navigate to the PasswordScreen
      // Navigate to the PasswordScreen with the phone number
      Navigator.pushNamed(
        context,
        '/password',
        arguments: PasswordScreenArguments(number.phoneNumber!),
      );
    }
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
              Image.asset('assets/logo.png',
                  height: 60), // Replace with your logo
              const SizedBox(height: 20),
              Image.asset('assets/pages/phone.png',
                  height: 220), // Replace with your phone illustration
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
                    InternationalPhoneNumberInput(
                      onInputChanged: (PhoneNumber number) {
                        this.number = number;
                      },
                      selectorConfig: const SelectorConfig(
                        selectorType: PhoneInputSelectorType.DROPDOWN,
                      ),
                      initialValue: number,
                      textFieldController: TextEditingController(),
                      inputDecoration: const InputDecoration(
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
                    icon: Image.asset(
                        'assets/icons/facebook.png'), // Replace with your Facebook icon
                    onPressed: () {},
                  ),
                  IconButton(
                    icon: Image.asset(
                        'assets/icons/google.png'), // Replace with your Google icon
                    onPressed: () {},
                  ),
                  IconButton(
                    icon: Image.asset(
                        'assets/icons/apple.png'), // Replace with your Apple icon
                    onPressed: () {},
                  ),
                ],
              ),
              const SizedBox(height: 10),
              // Sign Up link
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
