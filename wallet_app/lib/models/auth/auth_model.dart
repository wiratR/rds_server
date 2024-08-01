class RegisterInput {
  final String firstName;
  final String lastName;
  final String userName;
  final String phone;
  final String email;
  final String password;
  final String passwordConfirm;

  RegisterInput({
    required this.firstName,
    required this.lastName,
    required this.userName,
    required this.phone,
    required this.email,
    required this.password,
    required this.passwordConfirm,
  });

  Map<String, dynamic> toJson() {
    return {
      'firstName': firstName,
      'lastName': lastName,
      'userName': userName,
      'phone': phone,
      'email': email,
      'password': password,
      'passwordConfirm': passwordConfirm,
    };
  }
}

class LoginInput {
  final String email;
  final String password;

  LoginInput({
    required this.email,
    required this.password,
  });

  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'password': password,
    };
  }
}

class LoginInputByPhone {
  final String phone;
  final String password;

  LoginInputByPhone({
    required this.phone,
    required this.password,
  });

  Map<String, dynamic> toJson() {
    return {
      'phone': phone,
      'password': password,
    };
  }
}
