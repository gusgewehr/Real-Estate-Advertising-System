import 'package:frontend/domain/entity/address.dart';

class AddressModel extends AddressEntity {

  AddressModel({
    required super.zipCode,
    required super.street,
    required super.complement,
    required super.city,
    required super.neighborhood,
    required super.stateAbbr,
  });

  factory AddressModel.fromJson(Map<String, dynamic> json) {
    return AddressModel(
        zipCode: json['zip_code'] ?? "",
        street: json['street'] ?? "",
        complement: json['complement'] ?? "",
        neighborhood: json['neighborhood'] ?? "",
        city: json['city'] ?? "",
        stateAbbr: json['state_abbr'] ?? "",
    );
  }


  Map<String, dynamic> toJson() {
    return {
      'zip_code': super.zipCode,
      'street': super.street,
      'complement': super.complement,
      'city': super.city,
      'neighborhood': super.neighborhood,
      'state_abbr': super.stateAbbr

    };

  }



  factory AddressModel.fromEntity(AddressEntity entity) {
    return AddressModel(
      zipCode: entity.zipCode,
      street: entity.street ,
      complement: entity.complement ,
      city: entity.city,
      neighborhood: entity.neighborhood,
      stateAbbr: entity.stateAbbr,
    );
  }


}