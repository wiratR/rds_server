// ignore_for_file: use_build_context_synchronously
import 'package:flutter/material.dart';
// ignore: depend_on_referenced_packages
import 'package:intl_phone_number_input/intl_phone_number_input.dart';
import '../../components/appbar/custom_app_bar.dart';
import '../../constants.dart';
import '../../service/auth/auth_service.dart';
import '../../utils/utils.dart';
import '../login/login_screen.dart';

class RegisterScreen extends StatefulWidget {
  const RegisterScreen({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _RegisterScreenState createState() => _RegisterScreenState();
}

class _RegisterScreenState extends State<RegisterScreen> {
  final AuthService authService = AuthService();
  final _formKey = GlobalKey<FormState>();
  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final TextEditingController _confirmPasswordController =
      TextEditingController();
  final TextEditingController _phoneController = TextEditingController();

  String initialCountry = 'TH';
  PhoneNumber number = PhoneNumber(isoCode: 'TH');
  bool _isAccepted = false;
  bool _isPasswordVisible = false;
  bool _isConfirmPasswordVisible = false;
  bool _isTermsAcceptedValid = true;

  Future<void> _createAccount() async {
    setState(() {
      _isTermsAcceptedValid = _isAccepted;
    });

    if (_formKey.currentState!.validate() && _isAccepted) {
      // Perform registration logic
      debugPrint('Name: ${_nameController.text}');
      debugPrint('Email: ${_emailController.text}');
      debugPrint('Password: ${_passwordController.text}');
      debugPrint('Confirm Password: ${_confirmPasswordController.text}');
      debugPrint('Phone number: ${number.phoneNumber}');

      // String fullName = "John Doe";
      List<String> fullname = splitName(_nameController.text);
      String firstName = fullname[0];
      String lastName = fullname[1];
      debugPrint('First Name: $firstName');
      debugPrint('Last Name: $lastName');

      String userName = extractUsername(_emailController.text);
      debugPrint('User Name: $userName');

      // Navigate to the Login page
      Navigator.pushNamed(context, '/');

      String? user = await authService.register(
        firstName,
        lastName,
        userName,
        _emailController.text,
        number.phoneNumber.toString(),
        _passwordController.text,
        _confirmPasswordController.text,
      );

      if (user != null) {
        // Registration successful
        debugPrint('Register successful! User: $user');
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Register successful. Please try login.'),
            duration: Duration(seconds: 3),
          ),
        );
        Navigator.pushReplacement(
          context,
          MaterialPageRoute(builder: (context) => LoginScreen()),
        );
      } else {
        // Registration failed
        debugPrint('Register failed');
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Register failed. Please try again.'),
            duration: Duration(seconds: 3),
          ),
        );
        Navigator.pushReplacement(
          context,
          MaterialPageRoute(builder: (context) => LoginScreen()),
        );
      }
    }
  }

  @override
  void dispose() {
    _nameController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    _confirmPasswordController.dispose();
    _phoneController.dispose();
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
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(height: 10.0),
              const Text(
                'Create Account',
                style: TextStyle(
                  fontSize: 24.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 16.0),
              Form(
                key: _formKey,
                child: Column(
                  children: <Widget>[
                    TextFormField(
                      controller: _nameController,
                      decoration: InputDecoration(
                        labelText: 'Name',
                        hintText: 'e.g. John Doe',
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8.0),
                        ),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter name';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16.0),
                    TextFormField(
                      controller: _emailController,
                      decoration: InputDecoration(
                        labelText: 'Email',
                        hintText: 'e.g. email@example.com',
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8.0),
                        ),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter email';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16.0),
                    InternationalPhoneNumberInput(
                      onInputChanged: (PhoneNumber number) {
                        this.number = number;
                      },
                      selectorConfig: const SelectorConfig(
                        selectorType: PhoneInputSelectorType.DROPDOWN,
                      ),
                      initialValue: number,
                      textFieldController: _phoneController,
                      inputDecoration: const InputDecoration(
                        labelText: 'Mobile number',
                        border: OutlineInputBorder(),
                      ),
                    ),
                    const SizedBox(height: 16.0),
                    TextFormField(
                      controller: _passwordController,
                      obscureText: !_isPasswordVisible,
                      decoration: InputDecoration(
                        labelText: 'Password',
                        hintText: 'Enter your password',
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8.0),
                        ),
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
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter password';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 16.0),
                    TextFormField(
                      controller: _confirmPasswordController,
                      obscureText: !_isConfirmPasswordVisible,
                      decoration: InputDecoration(
                        labelText: 'Confirm Password',
                        hintText: 'Enter your confirm password',
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8.0),
                        ),
                        suffixIcon: IconButton(
                          icon: Icon(
                            _isConfirmPasswordVisible
                                ? Icons.visibility
                                : Icons.visibility_off,
                          ),
                          onPressed: () {
                            setState(() {
                              _isConfirmPasswordVisible =
                                  !_isConfirmPasswordVisible;
                            });
                          },
                        ),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter confirm password';
                        } else if (value != _passwordController.text) {
                          return 'Password mismatch';
                        }
                        return null;
                      },
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 16.0),
              Row(
                children: [
                  Checkbox(
                    value: _isAccepted,
                    onChanged: (bool? value) {
                      setState(() {
                        _isAccepted = value ?? false;
                      });
                    },
                  ),
                  const Text('I accept '),
                  GestureDetector(
                    onTap: () {
                      // Handle terms and conditions click
                    },
                    child: const Text(
                      'terms and conditions',
                      style: TextStyle(color: Colors.blue),
                    ),
                  ),
                ],
              ),
              Row(
                children: [
                  const Padding(
                    padding: EdgeInsets.all(16.0),
                  ), //Padding
                  const Text('    and '),
                  GestureDetector(
                    onTap: () {
                      // Handle privacy policy click
                    },
                    child: const Text(
                      'privacy policy',
                      style: TextStyle(color: Colors.blue),
                    ),
                  ),
                ],
              ),
              if (!_isTermsAcceptedValid)
                const Padding(
                  padding: EdgeInsets.symmetric(vertical: 8.0),
                  child: Text(
                    'You must accept terms and conditions.',
                    style: TextStyle(color: Colors.red),
                  ),
                ),
              const SizedBox(height: 16.0),
              ElevatedButton(
                onPressed: _createAccount,
                style: ElevatedButton.styleFrom(
                  backgroundColor: kPrimaryColor,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  minimumSize: const Size(double.infinity, 48.0),
                  foregroundColor: Colors.white,
                ),
                child: const Text('Create a new account'),
              ),
              const SizedBox(height: 16.0),
              const Center(child: Text('or continue using')),
              const SizedBox(height: 16.0),
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
            ],
          ),
        ),
      ),
    );
  }
}
