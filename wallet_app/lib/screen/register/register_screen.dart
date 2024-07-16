import 'package:flutter/material.dart';
import '../../components/appbar/custom_app_bar.dart';

class RegisterScreen extends StatelessWidget {
  const RegisterScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png', // Path to your logo asset
        showBackButton:
            true, // Set this to true if you want to show the back button
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              SizedBox(height: 50.0),
              Text(
                'Create Account',
                style: TextStyle(
                  fontSize: 24.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 16.0),
              TextField(
                decoration: InputDecoration(
                  labelText: 'Name',
                  hintText: 'e.g. John Doe',
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                ),
              ),
              SizedBox(height: 16.0),
              TextField(
                decoration: InputDecoration(
                  labelText: 'Email',
                  hintText: 'e.g. email@example.com',
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                ),
              ),
              SizedBox(height: 16.0),
              TextField(
                obscureText: true,
                decoration: InputDecoration(
                  labelText: 'Password',
                  hintText: 'Enter your password',
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  suffixIcon: Icon(Icons.visibility_off),
                ),
              ),
              SizedBox(height: 16.0),
              TextField(
                obscureText: true,
                decoration: InputDecoration(
                  labelText: 'Confirm Password',
                  hintText: 'Enter your confirm password',
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  suffixIcon: Icon(Icons.visibility_off),
                ),
              ),
              SizedBox(height: 16.0),
              Row(
                children: [
                  Checkbox(value: false, onChanged: (bool? value) {}),
                  Text('I accept '),
                  GestureDetector(
                    onTap: () {
                      // Handle terms and conditions click
                    },
                    child: Text(
                      'terms and conditions',
                      style: TextStyle(color: Colors.blue),
                    ),
                  ),
                ],
              ),
              // SizedBox(height: 16.0),
              Row(
                children: [
                  Padding(
                    padding: EdgeInsets.all(16.0),
                  ), //Padding
                  Text('    and '),
                  GestureDetector(
                    onTap: () {
                      // Handle privacy policy click
                    },
                    child: Text(
                      'privacy policy',
                      style: TextStyle(color: Colors.blue),
                    ),
                  ),
                ],
              ),
              SizedBox(height: 16.0),
              ElevatedButton(
                onPressed: () {
                  // Handle create account
                },
                style: ElevatedButton.styleFrom(
                  backgroundColor: Color(0xFF6C63FF), // Replace with your color
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  minimumSize: Size(double.infinity, 48.0),
                  foregroundColor: Colors.white,
                ),
                child: Text('Create a new account'),
              ),
              SizedBox(height: 16.0),
              Center(child: Text('or continue using')),
              SizedBox(height: 16.0),
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
