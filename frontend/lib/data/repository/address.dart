import '../../domain/entity/address.dart';
import '../../domain/repository/address.dart';
import '../datasources/address.dart';

class AddressRepositoryImpl implements AddressRepository{
  final AddressDataSource addressDataSource;

  AddressRepositoryImpl({required this.addressDataSource});

  @override
  Future<AddressEntity> getAddress(String cep) async {
    final res = await addressDataSource.getAddress(cep);

    return res;
  }



}