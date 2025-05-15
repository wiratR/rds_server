class AccountCreate {
  final String accountType;
  final String userId;

  AccountCreate({
    required this.accountType,
    required this.userId,
  });

  // Factory constructor to create an instance from JSON
  factory AccountCreate.fromJson(Map<String, dynamic> json) {
    return AccountCreate(
      accountType: json['account_type'],
      userId: json['user_id'],
    );
  }

  // Method to convert an instance to JSON
  Map<String, dynamic> toJson() {
    return {
      'account_type': accountType,
      'user_id': userId,
    };
  }
}
