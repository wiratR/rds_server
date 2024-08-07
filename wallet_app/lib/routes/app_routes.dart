import 'package:flutter/material.dart';
import '../screen/login/login_screen.dart';
import '../screen/login/otp_verification_screen.dart';
import '../screen/login/password_screen.dart';
import '../screen/main/home/home_screen.dart';
import '../screen/main/home/topup/card/payment_completion_screen.dart';
import '../screen/main/home/topup/card/payment_failed_screen.dart';
import '../screen/main/home/topup/topup_screen.dart';
import '../screen/register/register_screen.dart';

class AppRoutes {
  static Route<dynamic> getRoutes(RouteSettings settings) {
    switch (settings.name) {
      case '/':
        return MaterialPageRoute(
          builder: (_) => LoginScreen(),
        );
      case '/password':
        final args = settings.arguments as PasswordScreenArguments;
        return MaterialPageRoute(
          builder: (_) => PasswordScreen(phoneNumber: args.phoneNumber),
        );
      case '/register':
        return MaterialPageRoute(
          builder: (_) => const RegisterScreen(),
        );
      case '/otp-verification':
        final args = settings.arguments as OtpVerificationArguments;
        return MaterialPageRoute(
          builder: (_) => OtpVerificationScreen(phoneNumber: args.phoneNumber),
        );
      case '/home':
        return MaterialPageRoute(
          builder: (_) => HomeScreen(),
        );
      case '/top-up':
        return MaterialPageRoute(
          builder: (_) => TopUpScreen(),
        );
      case '/top-up/payment-complete':
        return MaterialPageRoute(
          builder: (_) => const PaymentCompletionScreen(
            paymentMessage: 'topup compleate',
          ),
        );
      case '/top-up/payment-failed':
        return MaterialPageRoute(
          builder: (_) => const PaymentFailedScreen(
            failureMessage: 'error topup',
          ),
        );
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

class OtpVerificationArguments {
  final String phoneNumber;

  OtpVerificationArguments(this.phoneNumber);
}

class PasswordScreenArguments {
  final String phoneNumber;

  PasswordScreenArguments(this.phoneNumber);
}
