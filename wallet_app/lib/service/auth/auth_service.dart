import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:shared_preferences/shared_preferences.dart';
// ignore: depend_on_referenced_packages
import 'package:http/http.dart' as http;
import 'package:wallet_app/models/user/user_model.dart';
import '../../config/app_config.dart';
import '../../models/auth/auth_model.dart';

class AuthService {
  Future<String?> login(String email, String password) async {
    // Using AppConfig.baseUrl
    final url = Uri.parse('${AppConfig.baseUrl}/auth/login');
    final headers = {'Content-Type': 'application/json'};
    final body =
        jsonEncode(LoginInput(email: email, password: password).toJson());

    debugPrint('Login with url = $url');
    debugPrint('Login with body = $body');

    try {
      final response = await http.post(url, headers: headers, body: body);

      if (response.statusCode == 200) {
        debugPrint('login successful');
        // Parse the JSON response
        Map<String, dynamic> jsonResponse = jsonDecode(response.body);
        // Access the token string
        String token = jsonResponse['data']['token'];
        // debugPrints the token as a string
        debugPrint('Token: $token');
        await saveToken(token);
        return token;
        // return response.body; // Example: return token or user data
      } else {
        debugPrint('login failed with status code: ${response.statusCode}');
        debugPrint('error = ${response.body}');
        return null;
      }
    } catch (e) {
      debugPrint('Error occurred: $e');
    }
    return null;
  }

  Future<String?> register(
      String firstName,
      String lastName,
      String userName,
      String phone,
      String email,
      String password,
      String passwordConfirm) async {
    final url = Uri.parse('${AppConfig.baseUrl}/auth/register');
    final headers = {'Content-Type': 'application/json'};
    final body = jsonEncode(RegisterInput(
            firstName: firstName,
            lastName: lastName,
            userName: userName,
            phone: phone,
            email: email,
            password: password,
            passwordConfirm: passwordConfirm)
        .toJson());

    debugPrint('Register with url = $url');
    debugPrint('Register with body = $body');

    try {
      final response = await http.post(url, headers: headers, body: body);

      if (response.statusCode == 201) {
        debugPrint('Register successful');
        return response.body; // Example: return token or user data
      } else {
        debugPrint('Register failed with status code: ${response.statusCode}');
        debugPrint('error = ${response.body}');
        return null;
      }
    } catch (e) {
      debugPrint('Error occurred: $e');
    }
    return null;
  }

  Future<String?> loginByPhone(String phone, String password) async {
    // Using AppConfig.baseUrl
    final url = Uri.parse('${AppConfig.baseUrl}/auth/loginbyphone');
    final headers = {'Content-Type': 'application/json'};
    final body = jsonEncode(
        LoginInputByPhone(phone: phone, password: password).toJson());

    debugPrint('Login with url = $url');
    debugPrint('Login with body = $body');

    try {
      final response = await http.post(url, headers: headers, body: body);

      if (response.statusCode == 200) {
        debugPrint('login successful');
        // Parse the JSON response
        Map<String, dynamic> jsonResponse = jsonDecode(response.body);
        // Access the token string
        String token = jsonResponse['data']['token'];
        // debugPrints the token as a string
        debugPrint('Token: $token');
        // Access the user string
        UserResponse userResponse =
            UserResponse.fromJson(jsonResponse['data']['user']);
        UserDetails user = userResponse.toUserDetails();
        // debugPrints the user id
        debugPrint('user id: ${user.id}');
        await saveToken(token);
        await saveUserId(user.id);
        return token;
        // return response.body; // Example: return token or user data
      } else {
        debugPrint('login failed with status code: ${response.statusCode}');
        debugPrint('error = ${response.body}');
        return null;
      }
    } catch (e) {
      debugPrint('Error occurred: $e');
    }
    return null;
  }

  Future<void> logout() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.remove('token');
    await prefs.remove('userid');
  }

  Future<bool> isAuthenticated() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String? token = prefs.getString('token');
    return token != null;
  }

  Future<String?> getToken() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.getString('token');
  }

  Future<void> saveToken(String token) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString('token', token);
  }

  Future<String?> getUserId() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.getString('userid');
  }

  Future<void> saveUserId(String token) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString('userid', token);
  }
}
