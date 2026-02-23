
import 'package:frontend/domain/entity/real_estate.dart';
import 'package:file_picker/file_picker.dart';

import '../entity/pagination.dart';
import '../repository/real_estate.dart';

class CreateRealEstateUseCase{
  final RealEstateRepository realEstateRepository;

  CreateRealEstateUseCase({required this.realEstateRepository});

  Future<void> createRealEstate(RealEstateEntity realEstate) async{
    await realEstateRepository.createRealEstate(realEstate);
  }
}

class GetRealEstateUseCase{
  final RealEstateRepository realEstateRepository;

  GetRealEstateUseCase({required this.realEstateRepository});

  Future<PaginationEntity<RealEstateEntity>> getRealEstate(int page, int pageSize) async{
    return await realEstateRepository.getRealEstate(page, pageSize);
  }
}

class UploadImageUseCase{
  final RealEstateRepository realEstateRepository;

  UploadImageUseCase({required this.realEstateRepository});

  Future<String> uploadImage(PlatformFile image) async{
    return await realEstateRepository.uploadImage(image);
  }
}