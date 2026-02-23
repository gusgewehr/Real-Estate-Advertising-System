
import '../entity/address.dart';
import '../repository/address.dart';

class GetAddressUseCase{
  final AddressRepository addressRepository;

  GetAddressUseCase({required this.addressRepository});

  Future<AddressEntity> getAddress(String cep) async{
    return await addressRepository.getAddress(cep);
  }


}