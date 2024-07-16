import 'package:flutter/material.dart';
import '../screen/login/login_screen.dart';
import '../screen/login/password_screen.dart';
import '../screen/register/register_screen.dart';

class AppRoutes {
  static Route<dynamic> getRoutes(RouteSettings settings) {
    switch (settings.name) {
      case '/':
        return MaterialPageRoute(builder: (_) => LoginScreen());
      case '/password':
        return MaterialPageRoute(builder: (_) => PasswordScreen());
      case '/register':
        return MaterialPageRoute(builder: (_) => RegisterScreen());
      // case '/home':
      //   return MaterialPageRoute(builder: (_) => HomeScreen());
      // Handle undefined routes
      default:
        return MaterialPageRoute(
          builder: (_) => Scaffold(
            body: Center(
              child: Text('No route defined for ${settings.name}'),
            ),
          ),
        );
    }
  }
}
