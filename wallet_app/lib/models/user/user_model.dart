class UserDetails {
  final String id;
  final String userName;
  final String firstName;
  final String lastName;
  final String email;
  final String phone;
  final String role;

  UserDetails({
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
      id: json['id'] as String,
      userName: json['userName'] as String,
      firstName: json['firstName'] as String,
      lastName: json['lastName'] as String,
      email: json['email'] as String,
      phone: json['phone'] as String,
      role: json['role'] as String,
    );
  }

  // Convert a UserDetails to a Map
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'userName': userName,
      'firstName': firstName,
      'lastName': lastName,
      'email': email,
      'phone': phone,
      'role': role,
    };
  }

  factory UserDetails.fromUserResponse(UserResponse userResponse) {
    return UserDetails(
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
      id: json['id'] as String,
      userName: json['userName'] as String,
      firstName: json['firstName'] as String,
      lastName: json['lastName'] as String,
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
      'id': id,
      'userName': userName,
      'firstName': firstName,
      'lastName': lastName,
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
