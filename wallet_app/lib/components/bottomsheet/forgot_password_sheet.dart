import 'package:flutter/material.dart';

class ForgotPasswordSheet extends StatelessWidget {
  const ForgotPasswordSheet({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              Text(
                'Forgot your password?',
                style: TextStyle(
                  fontSize: 18.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              TextButton(
                onPressed: () => Navigator.of(context).pop(),
                child: Text(
                  'Done',
                  style: TextStyle(
                    color: Colors.blue,
                  ),
                ),
              ),
            ],
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
          ElevatedButton(
            onPressed: () {
              // Handle send reset link
            },
            style: ElevatedButton.styleFrom(
              backgroundColor: Color(0xFF6C63FF), // Replace with your color
              foregroundColor: Colors.white,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(8.0),
              ),
              minimumSize: Size(double.infinity, 48.0), // Full-width button
            ),
            child: Text('Send reset link'),
          ),
          SizedBox(height: 16.0),
          Center(
            child: TextButton(
              onPressed: () {
                // Handle use mobile instead
              },
              child: Text('Use mobile instead'),
            ),
          ),
        ],
      ),
    );
  }
}
