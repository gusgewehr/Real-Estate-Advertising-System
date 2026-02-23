import 'package:frontend/domain/entity/address.dart';

class RealEstateEntity {
  final AddressEntity address;
  final double price;
  final String transactionType;
  final String imageUrl;

  RealEstateEntity({
    required this.address,
    required this.price,
    required this.transactionType,
    required this.imageUrl,
  });
}