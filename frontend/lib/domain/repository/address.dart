import 'package:frontend/domain/entity/address.dart';

abstract class AddressRepository {
  Future<AddressEntity> getAddress(String cep);
}