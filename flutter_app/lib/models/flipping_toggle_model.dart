class FlippingToggle {
  final bool enabled;

  FlippingToggle({required this.enabled});

  factory FlippingToggle.fromJson(Map<String, dynamic> json) {
    return FlippingToggle(
      enabled: json['enabled'] ?? false,
    );
  }

  Map<String, dynamic> toJson() => {
    'enabled': enabled,
  };
}
