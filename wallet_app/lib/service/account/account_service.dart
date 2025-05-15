// create account
// curl -X 'POST' \
//   'http://127.0.0.1:8000/api/accounts/create' \
//   -H 'accept: application/json' \
//   -H 'Content-Type: application/json' \
//   -d '{
//   "account_type": "mobile",
//   "user_id": "b1342091-76b5-4d04-b150-3b061d4964ad"
// }'
// Response
// {
//   "status": "success",
//   "data": {
//     "account": {
//       "id": "e8625956-4416-478b-9017-d2a027aae769",
//       "account_token": "5A7C47A3-CEC5-42E3-8B1D-57D78DF56224",
//       "account_type": "mobile",
//       "status": "bining",
//       "last_entry_time": "0001-01-01T00:00:00Z",
//       "active": true,
//       "txn_histories_detail": [
//         {
//           "txn_ref_id": "5A7C47A3-CEC5-42E3-8B1D-57D78DF56224",
//           "txn_detail": {
//             "txn_type_id": 1,
//             "txn_type_name": "Binding"
//           },
//           "sp_detail": {
//             "sp_id": 9
//           },
//           "loc_entry_detail": {
//             "line_detail": {
//               "line_id": -1
//             }
//           },
//           "loc_exit_detail": {
//             "loc_id": 99,
//             "loc_name": "RTV service",
//             "line_detail": {
//               "line_name": "Line 0"
//             }
//           },
//           "equipment_number": "mobile",
//           "created_at": "2024-09-12T00:44:18.769413Z",
//           "updated_at": "2024-09-12T00:44:18.769413Z"
//         }
//       ],
//       "created_at": "2024-09-12T00:44:18.750103343Z",
//       "updated_at": "2024-09-12T00:44:18.750103343Z"
//     }
//   }
// }

import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import '../../config/app_config.dart';
import '../../models/account/account_model.dart';

class AccountService {
  final String _authToken; // Add an instance variable for the auth token

  AccountService(this._authToken); // Constructor to accept the token

  Future<String?> createAccount(String accountType, String userId) async {
    final url = Uri.parse('${AppConfig.baseUrl}/accounts/create');
    final headers = {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $_authToken', // Add the authorization header
    };

    final body = jsonEncode(
        AccountCreate(accountType: accountType, userId: userId).toJson());

    debugPrint(url.toString());
    debugPrint(headers.toString());

    try {
      final response = await http.post(url, headers: headers, body: body);

      if (response.statusCode == 201) {
        // Parse the JSON response
        // Map<String, dynamic> jsonResponse = jsonDecode(response.body);
        return response.body; // Example: return token or user data
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
}
