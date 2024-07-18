class Member {
  final int id;
  final String first_name;
  final String last_name;
  final String username;
  final String? photo;
  final String email;
  final String? password;
  final String address;

  Member({
    required this.id,
    this.photo,
    required this.first_name,
    required this.last_name,
    required this.username,
    required this.email,
    this.password,
    required this.address,
  });

  factory Member.fromJson(Map<String, dynamic> json) {
    return Member(
      id: json['id'],
      first_name: json['first_name'],
      last_name: json['last_name'],
      username: json['username'],
      photo: json['photo'],
      email: json['email'],
      password: json['password'],
      address: json['address'],
    );
  }
}
