class UserDetails {
  final String accountId;
  final String id;
  final String userName;
  final String firstName;
  final String lastName;
  final String email;
  final String phone;
  final String role;

  UserDetails({
    required this.accountId,
    required this.id,
    required this.userName,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phone,
    required this.role,
  });

  // Convert a UserDetails from a Map
  factory UserDetails.fromJson(Map<String, dynamic> json) {
    return UserDetails(
      accountId: json['account_id'] as String,
      id: json['id'] as String,
      userName: json['user_name'] as String,
      firstName: json['first_name'] as String,
      lastName: json['last_name'] as String,
      email: json['email'] as String,
      phone: json['phone'] as String,
      role: json['role'] as String,
    );
  }

  // Convert a UserDetails to a Map
  Map<String, dynamic> toJson() {
    return {
      'account_id': accountId,
      'id': id,
      'user_name': userName,
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
      'phone': phone,
      'role': role,
    };
  }

  factory UserDetails.fromUserResponse(UserResponse userResponse) {
    return UserDetails(
      accountId: userResponse.accountId,
      id: userResponse.id,
      userName: userResponse.userName,
      firstName: userResponse.firstName,
      lastName: userResponse.lastName,
      email: userResponse.email,
      phone: userResponse.phone,
      role: userResponse.role,
    );
  }
}

class UserResponse {
  final String accountId;
  final String id;
  final String userName;
  final String firstName;
  final String lastName;
  final String email;
  final String phone;
  final String role;
  final DateTime createdAt;
  final DateTime updatedAt;

  UserResponse({
    required this.accountId,
    required this.id,
    required this.userName,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phone,
    required this.role,
    required this.createdAt,
    required this.updatedAt,
  });

  // Convert a UserResponse from a Map
  factory UserResponse.fromJson(Map<String, dynamic> json) {
    return UserResponse(
      accountId: json['account_id'] as String,
      id: json['id'] as String,
      userName: json['user_name'] as String,
      firstName: json['first_name'] as String,
      lastName: json['last_name'] as String,
      email: json['email'] as String,
      phone: json['phone'] as String,
      role: json['role'] as String,
      createdAt: DateTime.parse(json['created_at'] as String),
      updatedAt: DateTime.parse(json['updated_at'] as String),
    );
  }

  // Convert a UserResponse to a Map
  Map<String, dynamic> toJson() {
    return {
      'account_id': accountId,
      'id': id,
      'user_name': userName,
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
      'phone': phone,
      'role': role,
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
    };
  }

  // Convert UserResponse to UserDetails
  UserDetails toUserDetails() {
    return UserDetails(
      accountId: accountId,
      id: id,
      userName: userName,
      firstName: firstName,
      lastName: lastName,
      email: email,
      phone: phone,
      role: role,
    );
  }
}
