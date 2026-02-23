import 'package:flutter/foundation.dart';
import 'package:frontend/domain/entity/real_estate.dart';

import 'address.dart';

class RealEstateModel extends RealEstateEntity {

  RealEstateModel({
    required super.address,
    required super.price,
    required super.transactionType,
    required super.imageUrl
});

  factory RealEstateModel.fromJSON(Map<String, dynamic> json) {

    return RealEstateModel(
      address: AddressModel.fromJson(json['address']),
      price: (json['value'] ?? 0).toDouble(),
      transactionType: json['type'] ?? "",
      imageUrl: json['image'] ?? "",

    );
  }

  Map<String, dynamic> toJSON() {
    return {
      "address": AddressModel.fromEntity(address).toJson(),
      "value": price,
      "type": transactionType,
      "image": imageUrl,
    };
  }

  factory RealEstateModel.fromEntity(RealEstateEntity entity) {
    return RealEstateModel(
      address: AddressModel.fromEntity(entity.address),
      price: entity.price,
      transactionType: entity.transactionType,
      imageUrl: entity.imageUrl,
    );
  }
}