import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:uuid/uuid.dart';
import '../../config/app_config.dart';
import '../../models/user/user_model.dart';

class UserService {
  final String _authToken; // Add an instance variable for the auth token

  UserService(this._authToken); // Constructor to accept the token

  Future<UserResponse?> getCurrentUser() async {
    final url = Uri.parse('${AppConfig.baseUrl}/users/me');
    final headers = {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $_authToken', // Add the authorization header
    };

    debugPrint(url.toString());
    debugPrint(headers.toString());

    try {
      final response = await http.get(url, headers: headers);
      debugPrint(response.statusCode.toString());
      debugPrint(response.body);

      if (response.statusCode == 200) {
        return UserResponse.fromJson(jsonDecode(response.body));
      } else {
        throw Exception('Failed to load user');
      }
    } catch (e) {
      debugPrint('Error occurred: $e');
    }
    return null;
  }

  Future<UserResponse?> getUserById(String uuid) async {
    final url = Uri.parse('${AppConfig.baseUrl}/users/$uuid');
    final headers = {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $_authToken', // Add the authorization header
    };
    // debugPrint(url.toString());
    // debugPrint(headers.toString());
    try {
      final response = await http.get(url, headers: headers);
      if (response.statusCode == 200) {
        debugPrint('getUserById successful');
        // Parse the JSON response
        Map<String, dynamic> jsonResponse = jsonDecode(response.body);
        // Access the user string
        UserResponse userResponse =
            UserResponse.fromJson(jsonResponse['data']['user']);
        return userResponse;
      } else {
        debugPrint(response.statusCode.toString());
        debugPrint(response.body);
        throw Exception('Failed to load user');
      }
    } catch (e) {
      debugPrint('Error occurred: $e');
    }
    return null;
  }
}
